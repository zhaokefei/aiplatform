package api

import (
	"reflect"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/zhaokefei/aiplatform/storage"
	"github.com/zhaokefei/aiplatform/types"
)


func getUserByContext(c *gin.Context) *storage.User {
	user, ok := c.Get("user")
	if !ok {
		return nil
	}
	return user.(*storage.User)
}

func structToMap(s interface{}) map[string]interface{} {
	m := make(map[string]interface{})
	v := reflect.ValueOf(s)
	t := reflect.TypeOf(s)
	for i := 0; i < t.NumField(); i++ {
		m[t.Field(i).Name] = v.Field(i).Interface()
	}
	return m
}


func HandleAppsInfo(c *gin.Context) {
	user := getUserByContext(c)
	apps, err := storage.GetApps(user.Role)
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	} 
	c.JSON(200, apps)
}


func HandleCreateApp(c *gin.Context) {
	var Body types.App
	err := c.ShouldBindJSON(&Body)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	_, err = storage.CreateApp(structToMap(Body))
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"status": true,
	})
}


func HandleAppInfo(c *gin.Context) {
	idParam := c.Param("app_id")
	appID, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	app, err := storage.GetAppByID(appID)
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(200, app)
}


func HandleUpdateApp(c *gin.Context) {
	var Body types.App
	err := c.ShouldBindJSON(&Body)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}	
	idParam := c.Param("app_id")
	appID, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	appOps, err := storage.NewAppOps(appID)
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}
	appOps.Update(structToMap(Body))
	c.JSON(200, gin.H{
		"status": true,
	})
}


func HandleDeleteApp(c *gin.Context) {
	idParam := c.Param("app_id")
	appID, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	appOps, err := storage.NewAppOps(appID)
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}
	appOps.Delete()
	c.JSON(200, gin.H{
		"status": true,
	})
}