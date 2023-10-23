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
		}

		if sessionID == "" || !storage.UserLogined(sessionID) {
			c.JSON(401, gin.H{
				"error": "session过期",
			})
			c.Abort()
		}

		// before request
		c.Next()

		// after request

	}
}