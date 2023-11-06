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
		// 登陆接口
		auth.POST("/logined", HandleLogined)
		auth.POST("/logout", HandleLogout)
		// 用户相关
		auth.POST("/roles", HandleRoles)
		auth.GET("/user", HandleUserInfo)
	}


	// admin用户操作接口
	adminAuth := basic.Group("admin", middleware.AdminAuth())
	{
		// 用户与权限设置
		adminAuth.POST("/user/role", HandleUserRole)
		adminAuth.POST("/users", HandleUsers)
		// 应用设置
		adminAuth.GET("/apps", HandleAppsInfo)
		adminAuth.POST("/apps", HandleCreateApp)
		adminAuth.GET("/apps/:app_id", HandleAppInfo)
		adminAuth.PUT("/apps/:app_id", HandleUpdateApp)
		adminAuth.DELETE("/apps/:app_id", HandleDeleteApp)
	}
	return r
}