package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Customer struct {
	// ID    string `form:"id" json:"id" binding:"required"`
	// Level string `form:"level" json:"level" binding:"required"`
	ID    string `form:"id" json:"id"`
	Level string `form:"level" json:"level"`
}

func main() {
	route := gin.Default()
	route.GET("/testing", test)
	route.POST("/testing", test)
	route.Run()
}

func test(c *gin.Context) {
	var customer Customer
	// //Bind Query String
	// if err := c.ShouldBindQuery(&customer); err != nil {
	// 	fmt.Println("ShouldBindQuery fault", err)
	// }
	//Bind Post Body
	if err := c.ShouldBind(&customer); err != nil {
		// fmt.Println("ShouldBind fault", err)
	}
	fmt.Printf("customer:%+v", customer)
	c.String(http.StatusOK, "OK")
}

func post(c *gin.Context) {
	customer := Customer{}
	if err := c.ShouldBind(&customer); err != nil {
		// switch err {
		// case 'Customer.Level' Error:Field validation for 'Level' failed on the 'required' tag:
		// fmt.Println(123)
		// default :
		// }
	}
}
