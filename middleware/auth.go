package middleware

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"github.com/zhaokefei/aiplatform/storage"
)


func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 校验session-id正确性
		sessionID, err := c.Cookie("session-id")
		fmt.Println("sessionID: ", sessionID)
		if err != nil {
			c.JSON(400, gin.H{
				"error": "session-id doesn't exist",
			})
			c.Abort()
			return
		}

		if sessionID == "" || !storage.UserLogined(sessionID) {
			c.JSON(401, gin.H{
				"error": "session过期",
			})
			c.Abort()
			return
		}

		user, err := storage.GetUserBySessionID(sessionID)
		if err != nil {
			c.JSON(401, gin.H{
				"error": "session过期",
			})
			c.Abort()
			return
		}
		c.Set("user", user)
		// before request
		c.Next()

		// after request

	}
}


func AdminAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		var user *storage.User
		userInterface, ok := c.Get("user")
		if !ok {
			// session-id
			sessionID, err := c.Cookie("session-id")
			if err != nil {
				c.JSON(400, gin.H{
					"error": "session-id doesn't exist",
				})
				c.Abort()
				return
			}
			// 判断是不是admin用户
			user, err = storage.GetUserBySessionID(sessionID)
			if err != nil {
				c.JSON(401, gin.H{
					"error": "session过期",
				})
				c.Abort()
				return
			}
		} else {
			user, ok = userInterface.(*storage.User)
			if !ok || user == nil {
				c.JSON(401, gin.H{
					"error": "权限不足",
				})
				c.Abort()
				return
			} 
		}
		fmt.Println("user: ", user.Role)
		if user.Role < storage.SystemAdminID {
			c.JSON(401, gin.H{
				"error": "权限不足",
			})
			c.Abort()
			return
		}
		c.Next()

	}
}