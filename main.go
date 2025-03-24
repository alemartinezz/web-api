// Start of file: main.go

package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// Create a default Gin engine (with Logger & Recovery middleware)
	r := gin.Default()

	// Add our custom error-handling middleware
	r.Use(CustomErrorMiddleware())

	// "ping" endpoint => returns a unified success response
	r.GET("/ping", func(c *gin.Context) {
		JSONSuccess(c, http.StatusOK, gin.H{"message": "pong"})
	})

	// "error" endpoint => forcibly triggers an error, to demonstrate unified error response
	r.GET("/error", func(c *gin.Context) {
		_ = c.AbortWithError(http.StatusBadRequest, fmt.Errorf("Something went wrong."))
	})

	// Start the server on localhost:8080
	r.Run(":8080")
}

// End of file: main.go
