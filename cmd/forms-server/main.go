package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	fmt.Println("Starting http server on 8080...")
	r.Run() // listen and serve on 0.0.0.0:8080
	fmt.Println("Server ready.")
}
