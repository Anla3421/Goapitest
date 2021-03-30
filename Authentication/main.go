package main

import (
	"fmt"
	"net/http"
	"server/Api"

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
	listenPort := ":8001"

	// subRouterAuthenticated := router.Group("/api/v1/PersonId", gin.BasicAuth(gin.Accounts{
	// 	"admin": "admin",
	// }))
	// subRouterAuthenticated.Any("/:IdValue", PostMethod)

	// router.GET("/api/hello", func(c *gin.Context) {
	// 	name := c.Param("name")
	// 	c.JSON(200, gin.H{
	// 		"message": "hello " + name,
	// 	})
	// })

	// router.GET("/api/test", func(c *gin.Context) {
	// 	c.JSON(200, Sql.Users())
	// })

	// router.GET("/api/testgorm", func(c *gin.Context) {
	// 	c.JSON(200, Gorm.Users())
	// })

	// router.POST("/api/actionlog", func(c *gin.Context) {
	// 	c.JSON(200, gin.H{
	// 		"status": "200",
	// 		"msg":    "success",
	// 		"result": gin.H{
	// 			"actionlog": Gorm.Actionlogapi(),
	// 		},
	// 	})
	// })

	router.POST("/api/actionlog", Api.Actionlog)

	router.POST("/api/form_post", func(c *gin.Context) {
		message := c.PostForm("message")
		fmt.Println(c.Request.URL)
		fmt.Println(c.Request.PostForm)
		// fmt.Println(c.PostForm)
		nick := c.DefaultPostForm("nick", "anonymous")
		c.JSON(200, gin.H{
			"status":  "posted",
			"message": message,
			"nick":    nick,
		})
	})

	router.Run(listenPort)
}
