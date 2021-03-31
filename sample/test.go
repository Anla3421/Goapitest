package maintest

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/pkg/errors"
	"gorm.io/gorm/logger"
)

type Customer2 struct {
	ID    string `form:"id" json:"id"`
	Level string `form:"level" json:"level"`
}

func mainbb() {
	route := gin.Default()
	route.GET("/testing", test2)
	route.POST("/testing", test2)
	route.Run()
}

func test2(c *gin.Context) {
	var customer Customer2

	if err := c.ShouldBindBodyWith(&customer, binding.JSON); err != nil {
		fmt.Println("ShouldBindJSON fault", err)
		response := errorhandler.NewResponse(errorCode.DecodeJsonError)
		c.JSON(200, response)
		response.SetExtra(err)
		logger.Error(response)
	}

	fmt.Printf("customer:%+v", customer)
	c.String(http.StatusOK, "OK")
}

func ErrorHandle() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		err := c.Errors.Last()
		if err == nil {
			return
		}

		// Use reflect.TypeOf(err.Err) to known the type of your error
		if error, ok := errors.Cause(err.Err).(*myspace.KindOfClientError); ok {
			c.JSON(400, gin.H{
				"error": "Blah blahhh",
			})
			return
		}
	}
}
