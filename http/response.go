package http

// Response is the response for authenticate
type Response struct {
	Version    string
	StatusCode int
	StatusText string
	Headers    map[string]string
	Body       []byte
}
