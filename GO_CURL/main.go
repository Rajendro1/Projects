package main

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)



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

func main() {
	HandleRoute()
}

