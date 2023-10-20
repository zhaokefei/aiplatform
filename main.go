package main

import (
	"github.com/gin-gonic/gin"

	"github.com/zhaokefei/aiplatform/api"
)

func main() {
	r := gin.Default()

	r.POST("/register", api.HandleRegister)
	r.POST("/login", api.HandleLogin)
	r.POST("/logined", api.HandleLogined)
	r.POST("/logout", api.HandleLogout)

	r.Run() // listen and serve on 0.0.0.0:8080
}
