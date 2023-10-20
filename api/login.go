package api

import (
	"errors"

	"github.com/gin-gonic/gin"

	"github.com/zhaokefei/aiplatform/auth"
	"github.com/zhaokefei/aiplatform/types"
)

func HandleLogin(c *gin.Context) {
	var body types.LoginBody
	err := c.BindJSON(&body)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	user, err := auth.NewUser(body.Username)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "username doesn't exist",
		})
		return
	}
	token, err := user.Login(body.Password)
	if err != nil {
		c.JSON(500, gin.H{
			"error": "login failed: " + err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"session": token,
	})
}

func HandleRegister(c *gin.Context) {
	var Body types.RegisterBody
	err := c.BindJSON(&Body)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	_, err = auth.NewUser(Body.Username)
	if err == nil {
		c.JSON(500, gin.H{
			"error": "username already exists",
		})
		return
	}
	params := map[string]string{
		"sex": Body.Sex,
		"age": Body.Age,
	}
	_, err = auth.Register(Body.Username, Body.Password, Body.AgainPassword, Body.Email, params)
	if err != nil {
		c.JSON(500, gin.H{
			"error": "register failed: " + err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"status": true,
	})
}

func HandleLogined(c *gin.Context) {
	var Body types.LoginedBody
	err := c.BindJSON(&Body)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	user, err := auth.NewUser(Body.Username)
	if err != nil {
		c.JSON(500, gin.H{
			"error": "username doesn't exist",
		})
		return
	}
	status, err := user.IsLogin()
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

func HandleLogout(c *gin.Context) {
	var Body types.LoginedBody
	err := c.BindJSON(&Body)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	user, err := auth.NewUser(Body.Username)
	if err != nil {
		c.JSON(500, gin.H{
			"error": "username doesn't exist",
		})
		return
	}
	status, err := user.Logout()
	if err != nil {
		c.JSON(500, gin.H{
			"error": "check login failed",
		})
		return
	}
	c.JSON(200, gin.H{
		"status": status,
	})

}
