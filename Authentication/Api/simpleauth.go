package Api

import (
	"fmt"
	"net/http"

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
