package router

import (
	"app/controller"
	"app/middleware"

	"github.com/gin-gonic/gin"
)

func StartRouter() *gin.Engine {
	engine := gin.Default()

	engine.Use(middleware.SetUp)

	engine.POST("/signup", controller.CreateUser)
	engine.POST("/login", controller.LoginUser)
	engine.PATCH("/comment", controller.CreateComment)

	// jwtのミドルウエアを噛ませる
	engine.Use(middleware.AuthMiddleware.MiddlewareFunc())
	{
		engine.POST("/test", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"user": "pong",
			})
		})
	}

	return engine
}
