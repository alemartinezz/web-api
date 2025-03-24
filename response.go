// Start of file: response.go

package main

// ResponseFormat is our shared JSON schema for both success and errors.
type ResponseFormat struct {
	Status   string      `json:"status"`
	Code     int         `json:"code"`
	Data     interface{} `json:"data"`
	Messages []string    `json:"messages"`
}


// End of file: response.go
