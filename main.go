package main

import (
	"github.com/gin-gonic/gin"

	"github.com/zhaokefei/aiplatform/api"
	"github.com/zhaokefei/aiplatform/middleware"
)

func main() {
	r := gin.Default()

	r.POST("/register", api.HandleRegister)
	r.POST("/login", api.HandleLogin)

	basic := r.Group("/", middleware.Auth())

	auth := basic.Group("auth", )
	{
		auth.POST("/logined", api.HandleLogined)
		auth.POST("/logout", api.HandleLogout)
		auth.POST("/user/role", api.HandleUserRole)
	}


	r.Run() // listen and serve on 0.0.0.0:8080
}
