package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.ForceConsoleColor()
	r := gin.Default()
	r.GET("/hello/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.JSON(200, gin.H{
			"message": "hello " + name,
		})
	})
	r.GET("/", GetMethod)
	r.POST("/", PostMethod)
	listenPort := "4000"

	r.Run(":" + listenPort)
}

func PostMethod(c *gin.Context) {
	fmt.Println("\napi.go 'PostMethod' called")
	message := "PostMethod called"
	c.JSON(http.StatusOK, message)
}
func GetMethod(c *gin.Context) {
	fmt.Println("\napi.go 'GetMethod' called")
	message := "GetMethod called"
	c.JSON(http.StatusOK, message)
}
