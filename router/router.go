package router

import (
	"app/controller"
	"app/middleware"

	"github.com/gin-gonic/gin"
)

func StartRouter() *gin.Engine {
	engine := gin.Default()

	engine.Use(middleware.SetUp)

	engine.GET("/", controller.GetUsers)
	engine.POST("/", controller.PostUser)

	return engine
}
