package main

import (
	"Docker-LCS/functions"
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Welcome to the homepage!",
		})
	})

	r.POST("/api", func(c *gin.Context) {
		s := time.Now()

		var req Request
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"result":     "Some error occurred while parsing input. Ensure that input is a valid JSON in correct format. Contact admin for more details.",
				"time_taken": "0",
				"start_time": s.String(),
			})
			return
		}

		lengthOnly := req.Length_Only
		str1 := req.Str1
		str2 := req.Str2

		var LCS_length int
		var LCS string

		if lengthOnly {
			LCS_length = functions.LCS_length_only(str1, str2)
		} else {
			LCS_length, LCS = functions.LCS_with_string(str1, str2)
		}

		e := time.Now()
		elapsedTime := e.Sub(s).Seconds()

		var result gin.H
		if lengthOnly {
			result = gin.H{
				"time_taken": fmt.Sprintf("%.2f", elapsedTime),
				"start_time": s.String(),
				"end_time":   e.String(),
				"lcs_length": LCS_length,
			}
		} else {
			result = gin.H{
				"time_taken": fmt.Sprintf("%.2f", elapsedTime),
				"start_time": s.String(),
				"end_time":   e.String(),
				"lcs_length": LCS_length,
				"lcs":        LCS,
			}
		}

		c.JSON(http.StatusOK, result)

	})

	r.POST("/api/shutdown", func(c *gin.Context) {
		shutdownServer(c)
	})

	r.GET("/api/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"result": "success",
		})
	})

	if err := r.Run(":8080"); err != nil {
		log.Fatal("Failed to run server: ", err)
	}
}

func shutdownServer(c *gin.Context) {
	// Assume there's a condition to stop the server, such as receiving a signal or a scheduled task
	// Here's a simple example of stopping after a delay
	time.Sleep(5 * time.Second) // Assume waiting for 5 seconds before stopping the server

	// Create a context with a timeout of 5 seconds
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Call the server's Shutdown method to gracefully shutdown the server
	if err := c.Request.Context().Done(); err != nil {
		log.Fatalf("Shutdown error: %v", err)
	}

	r := gin.Default() // Declare the variable 'r' as a gin engine
	server := &http.Server{Addr: ":8080", Handler: r}
	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Server shutdown error: %v", err)
		c.JSON(http.StatusOK, gin.H{
			"result": fmt.Sprintf("Server shutdown error: %v", err),
		})
	}
	log.Println("Server stopped gracefully")
	c.JSON(http.StatusOK, gin.H{
		"result": "Server stopped gracefully",
	})
}

type Request struct {
	Length_Only bool   `json:"LengthOnly"`
	Str1        string `json:"str1"`
	Str2        string `json:"str2"`
}
