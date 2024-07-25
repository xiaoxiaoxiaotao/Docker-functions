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
			"result": "Welcome to the homepage!",
		})
	})

	r.POST("/api", func(c *gin.Context) {
		s := time.Now().UnixMicro()

		// Define an instance of the Request struct to store the parsing result
		var req Request
		if err := c.ShouldBindJSON(&req); err != nil {
			fmt.Println(err)
			c.JSON(http.StatusBadRequest, gin.H{
				"result":     "Some error occurred while parsing input. Ensure that input is a valid JSON in correct format. Contact admin for more details.",
				"time_taken": "0",
				"start_time": s,
			})
			return
		}

		lengthOnly := req.Root.Length_Only
		str1 := req.Root.Str1
		str2 := req.Root.Str2

		var LCS_length int
		var LCS string

		if lengthOnly {
			LCS_length = functions.LCS_length_only(str1, str2)
		} else {
			LCS_length, LCS = functions.LCS_with_string(str1, str2)
		}
		e := time.Now().UnixMicro()
		elapsedTime := e - s

		var result Result
		if lengthOnly {
			result = Result{
				LCS_length: LCS_length,
				LCS:        "default",
			}
		} else {
			result = Result{
				LCS_length: LCS_length,
				LCS:        LCS,
			}
		}

		response := Response{
			Result:     result,
			Time_taken: elapsedTime,
			Start_time: s,
		}
		fmt.Println("LCS Length: ", response)

		c.JSON(http.StatusOK, response)

	})

	r.POST("/api/shutdown", func(c *gin.Context) {
		shutdownServer(c)
	})

	r.GET("/api/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"result": "success",
		})
	})

	if err := r.Run(":5000"); err != nil {
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
	server := &http.Server{Addr: ":5000", Handler: r}
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
	Root RequestData `json:"root"`
}

type RequestData struct {
	Length_Only bool   `json:"LengthOnly"`
	Str1        string `json:"str1"`
	Str2        string `json:"str2"`
}

type Result struct {
	LCS_length int    `json:"LCS_length"`
	LCS        string `json:"LCS"`
}

type Response struct {
	Result     Result `json:"result"`
	Time_taken int64  `json:"time_taken"`
	Start_time int64  `json:"start_time"`
}
