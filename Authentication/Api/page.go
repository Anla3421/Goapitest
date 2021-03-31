package Api

import (
	"fmt"
	"server/Gorm"

	"github.com/gin-gonic/gin"
)

// type Request struct {
// 	Message string `json:message`
// 	ID      int    `json:id`
// }

func Actionlog(c *gin.Context) {
	// message := c.PostForm("message")
	// fmt.Println(c.Request.URL)
	// fmt.Println(c.Request.PostForm)
	// fmt.Println(c.Request)
	if c.Request.PostForm == nil {
		c.JSON(200, gin.H{
			"status": "400",
			"msg":    "bad request",
			// "message": message,
		})
		return
	}

	// All Green
	c.JSON(200, gin.H{
		"status": "200",
		"msg":    "success",
		"result": gin.H{
			"actionlog": Gorm.Actionlogapi(),
		},
	})

	fmt.Println(c.Request.URL.Query().Get("lastname"))
	// func NewRequest(method, url string, body io.Reader) (*Request, error)
	// func (r *Request) UserAgent() string
}
