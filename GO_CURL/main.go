package main

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
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

func ParseCurlCommand(curlCommand string) (ParsedCurl, error) {
	parsed := ParsedCurl{
		Method:     "GET", // Default method, will change to POST if data is provided
		Headers:    make(map[string]string),
		AuthParams: make(map[string]string), // For additional auth parameters
	}

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

// handleMethod handles the HTTP method parsing.
func handleMethod(parts []string, i int, parsed *ParsedCurl) int {
	if i+1 < len(parts) {
		parsed.Method = strings.ToUpper(parts[i+1])
		return i + 1
	}
	return i
}

func handleHeader(parts []string, i int, parsed *ParsedCurl) int {
	if i+1 < len(parts) {
		// The next part should be the full header, so capture it
		rawHeader := parts[i+1]

		// Check if the raw header part starts with 'Authorization:'
		if strings.HasPrefix(rawHeader, "'Authorization:") {
			// Combine the parts of the header if it's split due to spaces/newlines
			for i+2 < len(parts) && !strings.HasPrefix(parts[i+2], "--") {
				rawHeader += " " + parts[i+2]
				i++
			}
		}

		// Now safely trim the surrounding quotes, if any
		headerValue := strings.Trim(rawHeader, "'\"")

		// Split on the first colon to separate the key and value
		splitIndex := strings.Index(headerValue, ":")
		if splitIndex != -1 {
			key := strings.TrimSpace(headerValue[:splitIndex])
			value := strings.TrimSpace(headerValue[splitIndex+1:])

			parsed.Headers[key] = value
			handleAuthorizationHeader(key, value, parsed)
		} else {
			fmt.Printf("Invalid header format: %s\n", headerValue)
		}
		return i + 1
	}
	return i
}

func handleAuthorizationHeader(key, value string, parsed *ParsedCurl) {
	if strings.ToLower(key) == "authorization" {
		authParts := strings.Fields(value)
		if len(authParts) == 0 {
			// Log or handle the case where the Authorization header is empty
			fmt.Println("Warning: Authorization header is empty or malformed.")
			return
		}

		authType := strings.ToLower(authParts[0])

		switch authType {
		case "basic":
			if len(authParts) > 1 {
				parsed.AuthType = "basic"
				parsed.AuthToken = authParts[1]
			} else {
				fmt.Println("Warning: Basic authorization token is missing.")
			}
		case "bearer":
			if len(authParts) > 1 {
				parsed.AuthType = "bearer"
				parsed.AuthToken = authParts[1]
			} else {
				fmt.Println("Warning: Bearer token is missing.")
			}
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
			if len(authParts) > 1 {
				parsed.AuthType = authType
				parsed.AuthToken = authParts[1]
			} else {
				fmt.Printf("Warning: Unrecognized authorization type or token is missing: %s\n", authType)
			}
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

	// Measure the time taken for the request
	startTime := time.Now()

	// Execute the request
	resp, err := client.Do(req)
	timeTaken := time.Since(startTime).Milliseconds()

	response := make(map[string]interface{})

	if err != nil {
		// Check if the error is a timeout
		if os.IsTimeout(err) {
			response["error"] = "request timeout"
		} else {
			response["error"] = err.Error()
		}
		response["time_taken_ms"] = timeTaken
		return response, err
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// Determine response type (JSON, XML, etc.)
	contentType := resp.Header.Get("Content-Type")
	if strings.Contains(contentType, "application/json") {
		var jsonData interface{}
		if err = json.Unmarshal(body, &jsonData); err != nil {
			return nil, fmt.Errorf("failed to parse JSON response: %w", err)
		}
		response["body"] = jsonData
	} else if strings.Contains(contentType, "application/xml") || strings.Contains(contentType, "text/xml") {
		var xmlData interface{} // Define or use an appropriate XML structure
		if err = xml.Unmarshal(body, &xmlData); err != nil {
			return nil, fmt.Errorf("failed to parse XML response: %w", err)
		}
		response["body"] = xmlData
	} else {
		// Default to raw string response
		response["body"] = string(body)
	}

	response["status_code"] = resp.StatusCode
	response["time_taken_ms"] = timeTaken

	return response, nil
}

func main() {
	HandleRoute()
}
func HandleRoute() {
	gin.SetMode(gin.ReleaseMode)

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		singleLineCommand := preprocessCurlCommand(c.Request.FormValue("curl_url"))

		parsed, err := ParseCurlCommand(singleLineCommand)
		if err != nil {
			fmt.Println("Error parsing cURL command:", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		response, err := ExecuteCurlCommand(parsed)
		if err != nil {
			fmt.Println("Error executing cURL command:", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "response": response})
			return
		}
		c.JSON(http.StatusOK, gin.H{"response": response})
	})

	r.GET("/new", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "DON")
	})

	r.Run(":80")
}

// preprocessCurlCommand cleans up the cURL command string.
func preprocessCurlCommand(curlCommand string) string {
	// Remove unwanted characters and regularize spaces
	command := strings.ReplaceAll(curlCommand, `\\`, "")
	command = strings.ReplaceAll(command, `\`, "")
	command = strings.ReplaceAll(command, "\n", " ")
	command = strings.ReplaceAll(command, "\t", "")
	command = strings.ReplaceAll(command, `curl `, "")
	command = strings.ReplaceAll(command, `    `, "")
	command = strings.ReplaceAll(command, "\\\n", " ")
	return command
}
