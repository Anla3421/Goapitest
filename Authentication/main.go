package main

import (
	"fmt"
	"net/http"
	"server/Api"
	"server/Gorm"
	"server/Sql"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	listenPort := ":8080"

	//simple Authenticated with popup window
	subRouterAuthenticated := router.Group("/api/auth", gin.BasicAuth(gin.Accounts{
		"admin": "admin",
	}))
	subRouterAuthenticated.GET("/:IdValue", Api.PostMethod)

	router.GET("/api/hello", func(c *gin.Context) {
		name := c.Param("name")
		c.JSON(200, gin.H{
			"message": "hello " + name,
		})
	})

	router.GET("/api/test", func(c *gin.Context) {
		c.JSON(200, Sql.Users())
	})

	router.GET("/api/testgorm", func(c *gin.Context) {
		c.JSON(200, Gorm.Users())
	})

	router.POST("/api/actionlog", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "200",
			"msg":    "success",
			"result": gin.H{
				"actionlog": Gorm.Actionlogapi(),
			},
		})
	})

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
