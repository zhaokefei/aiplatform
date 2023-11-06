package types

import "github.com/zhaokefei/aiplatform/storage"


type App struct {
	Name string `json:"name"`
	Order int `json:"order"`
	Description string `json:"description"`
	Icon string `json:"icon"`
	Mode storage.AppMode `json:"mode"`
	Path string `json:"path"`
	RoleID int `json:"role_id"`
	CategoryID int `json:"category_id"`
}