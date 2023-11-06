package api

import (
	"github.com/gin-gonic/gin"
	"github.com/zhaokefei/aiplatform/storage"
)


func HandleUserInfo(c *gin.Context) {
	sessionID, _ := c.Cookie("session-id")
	user, err := storage.GetUserBySessionID(sessionID)
	if err != nil {
		c.JSON(500, gin.H{
			"error": "username doesn't exist",
		})
		return
	}
	c.JSON(200, user)
}


func HandleRoles(c *gin.Context) {
	sessionID, _ := c.Cookie("session-id")
	user, err := storage.GetUserBySessionID(sessionID)
	if err != nil {
		c.JSON(500, gin.H{
			"error": "username doesn't exist",
		})
		return
	}
	// 获取当前用户可看到权限的roles
	roles, err := storage.GetRoles(user.Role)
	if err != nil {
		c.JSON(500, gin.H{
			"error": "get roles failed: " + err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"roles": roles,
	})
}