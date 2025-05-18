package main

import (
	"fmt"
	"log"
	"net/http"

	gin "github.com/gin-gonic/gin"
)

func handler(c *gin.Context) {
	c.JSON(200, "Hello from Gin in Docker!")
}

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) { handler(c) })

	fmt.Println("Go backend started!")
	log.Fatal(http.ListenAndServe(":8080", r))
}
