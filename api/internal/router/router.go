package router

import (
	"go-gin-rpc/api/internal/controller/article"
	"go-gin-rpc/api/internal/controller/greeting"

	"github.com/gin-gonic/gin"
)

func SetApiRouter(router *gin.Engine) {
	// Simple group: v1
	group := router.Group("/api")
	{
		group.GET("/article/detail", article.Detail)

		group.GET("/greeting/:name", greeting.Hello)
	}
}
