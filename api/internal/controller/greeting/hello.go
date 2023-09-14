package greeting

import (
	"fmt"
	v1 "go-gin-rpc/gen/greeting/v1"
	"go-gin-rpc/rpc/client/greeting"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Hello(c *gin.Context) {
	name := c.Param("name")

	client, err := greeting.NewClient()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	// Contact the server and print out its response.
	req := &v1.HelloRequest{Name: name}
	res, err := client.SayHello(c, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"result": fmt.Sprint(res.Message),
	})

}
