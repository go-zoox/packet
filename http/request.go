package http

// Request is the request for authenticate
//
// Example:
// GET /get HTTP/1.1
// Host: localhost:8081
// Connection: keep-alive
// sec-ch-ua: "Google Chrome";v="89", "Chromium";v="89", ";Not A Brand";v="99"
// age: 10
// name: zhufeng
// sec-ch-ua-mobile: ?0
// User-Agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_5) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/89.0.4389.128 Safari/537.36
// Accept: */*
// Sec-Fetch-Site: same-origin
// Sec-Fetch-Mode: cors
// Sec-Fetch-Dest: empty
// Referer: http://localhost:8081/get.html
// Accept-Encoding: gzip, deflate, br
// Accept-Language: zh-CN,zh;q=0.9,en;q=0.8
//
type Request struct {
	Method  string
	URI     string
	Version string
	//
	Headers map[string]string
	//
	Body []byte
}
