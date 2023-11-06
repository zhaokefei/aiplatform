package main

import (
	"bytes"
	"encoding/json"
	"net/http/httptest"
	"testing"

	"github.com/zhaokefei/aiplatform/api"
)


func init() {
	router = api.Routers()
}


func loginGetCookie(t *testing.T, username string, password string) string {
	// 准备数据
	data := map[string]string{
		"username": username,
		"password": password,
	}
	requestBody, _ := json.Marshal(data)

	w := httptest.NewRecorder()
	req := httptest.NewRequest("POST", "/login", bytes.NewBuffer(requestBody))
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)

	if w.Code != 200 {
		t.Fatalf("handler returned wrong status code: got %v want %v", w.Code, 200)
	}

	cookies := w.Result().Cookies()
	if len(cookies) == 0 {
		t.Fatalf("handler not get cookies")
	}
	return cookies[0].Value

}


func TestAppRoute_Fail(t *testing.T) {
	// 准备数据
	// loginCookie := loginGetCookie(t, "admin", "admin")

}