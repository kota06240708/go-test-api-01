package router

import (
	"app/controller"
	"app/middleware"

	"github.com/gin-gonic/gin"
)

func StartRouter() *gin.Engine {
	engine := gin.Default()

	engine.Use(middleware.SetUp)

	engine.POST("/user", controller.CreateUser)
	engine.POST("/login", controller.LoginUser)
	engine.PATCH("/comment", controller.CreateComment)

	return engine
}
