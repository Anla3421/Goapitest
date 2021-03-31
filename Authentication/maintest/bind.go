package maintest

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type Customer struct {
	ID    string `form:"id" json:"id" binding:"required"`
	Level string `form:"level" json:"level" binding:"required"`
}

func mainaa() {
	route := gin.Default()
	route.GET("/testing", test)
	route.POST("/testing", test)
	route.Run()
}

func test(c *gin.Context) {
	var customer Customer
	//Bind Query String
	if err := c.ShouldBindQuery(&customer); err != nil {
		fmt.Println("ShouldBindQuery fault", err)
	}
	//Bind Post Body
	if err := c.ShouldBindBodyWith(&customer, binding.JSON); err != nil {
		fmt.Println("ShouldBindJSON fault", err)
	}
	fmt.Printf("customer:%+v", customer)
	c.String(http.StatusOK, "OK")
}
