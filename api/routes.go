package api

import (
	"github.com/gin-gonic/gin"

	"github.com/zhaokefei/aiplatform/middleware"
)


func Routers() *gin.Engine {
	r := gin.Default()

	r.POST("/register", HandleRegister)
	r.POST("/login", HandleLogin)

	basic := r.Group("/", middleware.Auth())

	// auth接口
	auth := basic.Group("auth")
	{
		auth.GET("/user", HandleUserInfo)
		auth.POST("/logined", HandleLogined)
		auth.POST("/logout", HandleLogout)
		auth.POST("/roles", HandleRoles)
	}

	// admin用户操作接口
	adminAuth := basic.Group("admin", middleware.AdminAuth())
	{
		adminAuth.POST("/user/role", HandleUserRole)
		adminAuth.POST("/users", HandleUsers)
	}
	return r
}