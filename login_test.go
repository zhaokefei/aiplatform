package main

import (
	"bytes"
	"encoding/json"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/zhaokefei/aiplatform/api"
)


var router *gin.Engine

func init() {
	router = api.Routers()
}


func TestLoginRoute_Fail(t *testing.T) {
	// 准备数据
	data := map[string]string{
		"username": "notsetuser",
		"password": "errorpassword",
	}
	requestBody, _ := json.Marshal(data)

	w := httptest.NewRecorder()
	req := httptest.NewRequest("POST", "/login", bytes.NewBuffer(requestBody))
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)

	if w.Code != 500 {
		t.Fatalf("handler returned wrong status code: got %v want %v", w.Code, 200)
	}
}


func TestLoginRoute_Success(t *testing.T) {
	// 准备数据
	data := map[string]string{
		"username": "admin",
		"password": "admin",
	}
	requestBody, _ := json.Marshal(data)

	w := httptest.NewRecorder()
	req := httptest.NewRequest("POST", "/login", bytes.NewBuffer(requestBody))
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)

	if w.Code != 200 {
		t.Fatalf("handler returned wrong status code: got %v want %v", w.Code, 200)
	}
}
