package main

import (
	"fmt"
	"net/http"
	"server/Sql"

	"github.com/gin-gonic/gin"
)

func PostMethod(c *gin.Context) {
	fmt.Println("\n'GetMethod' called")
	IdValue := c.Params.ByName("IdValue")
	message := "GetMethod Called With Param: " + IdValue
	c.JSON(http.StatusOK, message)
	ReqPayload := make([]byte, 1024)
	ReqPayload, err := c.GetRawData()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Request Payload Data: ", string(ReqPayload))
}

func main() {
	router := gin.Default()
	subRouterAuthenticated := router.Group("/api/v1/PersonId", gin.BasicAuth(gin.Accounts{
		"admin": "admin",
	}))
	subRouterAuthenticated.Any("/:IdValue", PostMethod)
	listenPort := ":8080"

	router.GET("/api/hello", func(c *gin.Context) {
		name := c.Param("name")
		c.JSON(200, gin.H{
			"message": "hello " + name,
		})
	})

	router.GET("/api/test", func(c *gin.Context) {
		c.JSON(200, Sql.Users())
	})

	router.Run(listenPort)
}
