// Start of file: errors.go
package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// CustomErrorMiddleware catches any errors set on c.Errors
// and turns them into a unified JSON response.
func CustomErrorMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Process all next handlers first
		c.Next()

		// Check if any errors were accumulated
		if len(c.Errors) > 0 {
			// You can pick the first error or handle multiple
			err := c.Errors[0].Err

			// If status was left as 200, override to 500
			statusCode := c.Writer.Status()
			if statusCode == http.StatusOK {
				statusCode = http.StatusInternalServerError
			}

			// Build the response shape
			resp := ResponseFormat{
				Status:   http.StatusText(statusCode),
				Code:     statusCode,
				Data:     nil,
				Messages: []string{err.Error()},
			}

			// Send unified error response
			c.JSON(statusCode, resp)
			c.Abort()
		}
	}
}


// End of file: errors.go
