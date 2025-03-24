// Start of file: success.go

package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// JSONSuccess sends a unified success response (same shape as errors).
func JSONSuccess(c *gin.Context, code int, data interface{}, messages ...string) {
	resp := ResponseFormat{
		Status:   http.StatusText(code),
		Code:     code,
		Data:     data,
		Messages: messages,
	}

	c.JSON(code, resp)
}


// End of file: success.go
