package article

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Detail(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Detail",
	})
}
