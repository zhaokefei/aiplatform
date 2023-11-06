package api

import (
	"errors"

	"github.com/gin-gonic/gin"

	"github.com/zhaokefei/aiplatform/auth"
	"github.com/zhaokefei/aiplatform/storage"
	"github.com/zhaokefei/aiplatform/types"
)

func HandleUserRole(c *gin.Context) {
	var Body types.UserRoleBody
	err := c.ShouldBindJSON(&Body)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	user, err := auth.NewUser(Body.Username, "")
	if err != nil {
		c.JSON(500, gin.H{
			"error": "username doesn't exist",
		})
		return
	}
	status, err := user.UserInfo.SetRole(Body.Rolename)
	if err != nil {
		c.JSON(500, gin.H{
			"error": errors.New("check login failed"),
		})
		return
	}
	c.JSON(200, gin.H{
		"status": status,
	})
}



func HandleUsers(c *gin.Context) {
	var body types.Users
	err := c.ShouldBindJSON(&body)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "invalid request body",
		})
		return
	}

	users, err := storage.GetUsers(body.Limit, body.Offset)
	if err != nil {
		c.JSON(500, gin.H{
			"error": "get users failed: " + err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"users": users,
	})
}
