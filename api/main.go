package main

import (
	"go-gin-rpc/api/internal/router"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func main() {

	// Set up a http server.
	engine := gin.Default()

	gin.SetMode(gin.DebugMode)
	// Switch to "release" mode in production
	// gin.SetMode(gin.ReleaseMode)

	// 跨域
	engine.Use(cors.Default())

	// 设置 API 路由
	router.SetApiRouter(engine)

	// Run http server
	if err := engine.Run(":8080"); err != nil {
		logrus.Fatalf("could not run server: %v", err)
	}
}
