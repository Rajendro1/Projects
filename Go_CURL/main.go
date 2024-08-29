package main

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

type ParsedCurl struct {
	Method       string
	Headers      map[string]string
	Body         string
	URL          string
	ContentType  string
	AuthType     string
	AuthToken    string
	AuthUsername string
	AuthPassword string
	AuthParams   map[string]string
	UserAgent    string
	DataType     string
}

// handleMethod handles the HTTP method parsing.
func handleMethod(parts []string, i int, parsed *ParsedCurl) int {
	if i+1 < len(parts) {
		parsed.Method = strings.ToUpper(parts[i+1])
		return i + 1
	}
	return i
}

// handleHeader handles the header parsing and identifies authorization types.
func handleHeader(parts []string, i int, parsed *ParsedCurl) int {
	if i+1 < len(parts) {
		header := strings.SplitN(parts[i+1], ":", 2)
		if len(header) == 2 {
			key := strings.TrimSpace(header[0])
			value := strings.TrimSpace(header[1])
			parsed.Headers[key] = value
			handleAuthorizationHeader(key, value, parsed)
		}
		return i + 1
	}
	return i
}

// handleAuthorizationHeader handles the various types of authorization headers.
func handleAuthorizationHeader(key, value string, parsed *ParsedCurl) {
	if strings.ToLower(key) == "authorization" {
		authParts := strings.Fields(value)
		authType := strings.ToLower(authParts[0])

		switch authType {
		case "basic":
			parsed.AuthType = "basic"
			parsed.AuthToken = authParts[1]
		case "bearer":
			parsed.AuthType = "bearer"
			parsed.AuthToken = authParts[1]
		case "digest":
			parsed.AuthType = "digest"
			parsed.AuthParams["digest"] = value
		case "oauth":
			handleOAuth(authParts, parsed)
		case "hawk":
			parsed.AuthType = "hawk"
			parsed.AuthParams["hawk"] = value
		case "aws":
			parsed.AuthType = "aws_signature"
			parsed.AuthParams["aws_signature"] = value
		case "ntlm":
			parsed.AuthType = "ntlm"
			parsed.AuthParams["ntlm"] = value
		case "akamai-edgegrid":
			parsed.AuthType = "akamai_edgegrid"
			parsed.AuthParams["akamai_edgegrid"] = value
		case "asap":
			parsed.AuthType = "asap"
			parsed.AuthParams["asap"] = value
		default:
			parsed.AuthType = authType
			parsed.AuthToken = authParts[1]
		}
	}
}

// handleOAuth handles specific OAuth types.
func handleOAuth(authParts []string, parsed *ParsedCurl) {
	if len(authParts) > 1 && authParts[1] == "1.0" {
		parsed.AuthType = "oauth1.0"
		parsed.AuthParams["oauth1.0"] = strings.Join(authParts, " ")
	} else if len(authParts) > 1 && authParts[1] == "2.0" {
		parsed.AuthType = "oauth2.0"
		parsed.AuthToken = authParts[2]
	}
}

// handleForm handles the form parsing and adjusts the method if necessary.
func handleForm(parts []string, i int, parsed *ParsedCurl) int {
	if i+1 < len(parts) {
		parsed.Body = parts[i+1]
		if parsed.Method == "GET" {
			parsed.Method = "POST"
		}
		if parsed.ContentType == "" {
			parsed.ContentType = "multipart/form-data"
		}
		return i + 1
	}
	return i
}

// handleBasicAuth handles basic authentication.
func handleBasicAuth(parts []string, i int, parsed *ParsedCurl) int {
	if i+1 < len(parts) {
		authParts := strings.SplitN(parts[i+1], ":", 2)
		if len(authParts) == 2 {
			parsed.AuthType = "basic"
			parsed.AuthUsername = authParts[0]
			parsed.AuthPassword = authParts[1]
		}
		return i + 1
	}
	return i
}

// handleUserAgent handles the user-agent parsing.
func handleUserAgent(parts []string, i int, parsed *ParsedCurl) int {
	if i+1 < len(parts) {
		parsed.UserAgent = parts[i+1]
		return i + 1
	}
	return i
}

// determineContentType determines the content type based on the body content.
func determineContentType(parsed *ParsedCurl) {
	if parsed.ContentType == "" && (strings.HasPrefix(parsed.Body, "{") || strings.HasPrefix(parsed.Body, "[")) {
		parsed.ContentType = "application/json"
	}

	if parsed.ContentType == "application/json" && strings.Contains(parsed.Body, "query") {
		parsed.ContentType = "application/graphql"
	}
}

// ExecuteCurlCommand executes the given cURL command and returns the response as an interface{}.
func ExecuteCurlCommand(parsed ParsedCurl) (map[string]interface{}, error) {
	// Create the HTTP client with a timeout
	client := &http.Client{Timeout: 10 * time.Second}
	// Create the HTTP request
	req, err := http.NewRequest(parsed.Method, parsed.URL, bytes.NewBuffer([]byte(parsed.Body)))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", parsed.ContentType)

	// Set headers
	for key, value := range parsed.Headers {
		req.Header.Set(key, value)
	}

	// Set basic auth if available
	if parsed.AuthType == "basic" {
		req.SetBasicAuth(parsed.AuthUsername, parsed.AuthPassword)
	}
	// Execute the request
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	// Determine response type (JSON, XML, etc.)
	contentType := resp.Header.Get("Content-Type")
	response := make(map[string]interface{})

	if strings.Contains(contentType, "application/json") {
		err = json.Unmarshal(body, &response)
		if err != nil {
			return nil, fmt.Errorf("failed to parse JSON response: %w", err)
		}
	} else if strings.Contains(contentType, "application/xml") || strings.Contains(contentType, "text/xml") {
		// Handle XML response
		err = xml.Unmarshal(body, &response)
		if err != nil {
			return nil, fmt.Errorf("failed to parse XML response: %w", err)
		}
	} else {
		// Default to raw string response
		response["body"] = string(body)
	}

	// Add the status code to the response map
	response["status_code"] = resp.StatusCode

	return response, nil
}
func extractJSONFromDataRaw(command string, dataType string) (string, error) {
	// Look for the start of the --data-raw argument
	startIndex := strings.Index(command, dataType)
	if startIndex == -1 {
		return "", fmt.Errorf("data-raw section not found")
	}

	// Move the start index forward to the beginning of the JSON data
	startIndex += len(dataType)

	// Extract the substring starting from startIndex to the end
	dataPart := command[startIndex:]

	// Trim leading whitespace and the first single quote
	dataPart = strings.TrimSpace(dataPart)
	if strings.HasPrefix(dataPart, "'") {
		dataPart = dataPart[1:]
	} else {
		return "", fmt.Errorf("data-raw section does not start with a quote as expected")
	}

	// Find the end of the JSON data, which should be the next single quote
	endIndex := strings.Index(dataPart, "'")
	if endIndex == -1 {
		return "", fmt.Errorf("closing quote for data-raw section not found")
	}

	// Extract the JSON part
	jsonData := dataPart[:endIndex]

	return jsonData, nil
}
func main() {
	gin.SetMode(gin.ReleaseMode)

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {

		singleLineCommand := c.Request.FormValue("data")
		singleLineCommand = strings.ReplaceAll(singleLineCommand, "\n", " ")
		singleLineCommand = strings.ReplaceAll(singleLineCommand, "\t", "") // Optional: remove tabs if necessary
		singleLineCommand = strings.ReplaceAll(singleLineCommand, `curl `, "")
		singleLineCommand = strings.ReplaceAll(singleLineCommand, `    `, "")
		singleLineCommand = strings.ReplaceAll(singleLineCommand, `\\`, "")
		singleLineCommand = strings.ReplaceAll(singleLineCommand, `\`, "")
		c.JSON(http.StatusOK, singleLineCommand)
		parsed, err := ParseCurlCommand(singleLineCommand)
		if err != nil {
			fmt.Println("Error parsing cURL command:", err)
			return
		}

		response, err := ExecuteCurlCommand(parsed)
		if err != nil {
			fmt.Println("Error executing cURL command:", err)
			return
		}
		c.JSON(http.StatusOK, gin.H{"response": response})

	})
	r.GET("/new", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "DON")
	})

	r.Run(":80")
}
func ParseCurlCommand(curlCommand string) (ParsedCurl, error) {
	parsed := ParsedCurl{
		Method:     "GET", // Default method, will change to POST if data is provided
		Headers:    make(map[string]string),
		AuthParams: make(map[string]string), // For additional auth parameters
	}

	// Remove line breaks to ensure the command is parsed as a single string
	curlCommand = strings.ReplaceAll(curlCommand, "\\\n", " ")

	// Split the command into parts
	parts := strings.FieldsFunc(curlCommand, func(c rune) bool {
		return c == ' '
	})

	// Process each part of the command
	for i := 0; i < len(parts); i++ {
		trimmedPart := strings.Trim(parts[i], "'\"") // Remove surrounding quotes

		if strings.HasPrefix(trimmedPart, "http://") || strings.HasPrefix(trimmedPart, "https://") {
			parsed.URL = trimmedPart
		}

		// Immediately after setting URL, check for data flags
		switch parts[i] {
		case "-X", "--request":
			i = handleMethod(parts, i, &parsed)
		case "-H", "--header":
			i = handleHeader(parts, i, &parsed)
		case "-d", "--data", "--data-urlencode", "--data-raw":
			parsed.DataType = parts[i]
			jsonRaw, err := extractJSONFromDataRaw(curlCommand, parsed.DataType)
			if err != nil {
				fmt.Println("Error parsing cURL command:", err)
			}
			parsed.Body = jsonRaw
			i = handleData(parts, i, &parsed)
		case "-F", "--form":
			i = handleForm(parts, i, &parsed)
		case "-u", "--user":
			i = handleBasicAuth(parts, i, &parsed)
		case "-A", "--user-agent":
			i = handleUserAgent(parts, i, &parsed)
		}
	}

	if parsed.URL == "" {
		return parsed, errors.New("no URL found")
	}

	determineContentType(&parsed)

	return parsed, nil
}

// handleData handles the data parsing and adjusts the method if necessary.
func handleData(parts []string, i int, parsed *ParsedCurl) int {
	if i+1 < len(parts) {
		// parsed.Body = parts[i+1]
		// If the method is still set to the default "GET", change it to "POST"
		if parsed.Method == "GET" {
			fmt.Println("Switching method to POST due to data flag") // Debugging: Show when the method is switched to POST
			parsed.Method = "POST"
		}

		return i + 1
	}
	return i
}
