package main

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
