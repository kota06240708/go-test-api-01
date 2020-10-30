package router

import (
	"app/controller"
	"app/middleware"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func StartRouter() *gin.Engine {
	engine := gin.Default()

	store := cookie.NewStore([]byte("secret"))
	engine.Use(sessions.Sessions("mysessions", store))

	engine.Use(middleware.SetUp)

	// クロスオリジン対応
	engine.Use(cors.New(cors.Config{
		AllowOrigins: []string{
			"http://localhost:8080",
			"http://localhost:8081",
			"http://localhost:8082",
			"http://localhost:8083",
			"http://localhost:3000",
			"https://localhost:3000",
			"https://dev-bylegal.jp",
			"https://kota06240708.postman.co/",
			"https://professional.attivita.co.jp",
			"https://service.attivita.co.jp",
			"https://storage.googleapis.com",
		},
		AllowMethods:     []string{"*"},
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		// AllowOriginFunc: func(origin string) bool {
		// 	return origin == "https://github.com"
		// },
		// MaxAge: 12 * time.Hour,
	}))

	engine.POST("/signup", controller.CreateUser)
	engine.POST("/login", controller.LoginUser)

	// jwtのミドルウエアを噛ませる
	engine.Use(middleware.AuthMiddleware.MiddlewareFunc())
	{
		engine.POST("/test", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"user": "pong",
			})
		})

		// トークンを再発行
		engine.PATCH("/refresh_token", controller.RefreshToken)

		// コメントをする
		engine.PATCH("/comment", controller.CreateComment)
	}

	return engine
}
