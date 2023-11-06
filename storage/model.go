package storage

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID        int `gorm:"primarykey"`
	Avatar    string `gorm:"not null;size:255"`
	Username  string `gorm:"uniqueindex;not null;size:255"`
	Password  string
	Age       string
	Sex       string
	Email     string `gorm:"uniqueindex;not null;size:255"`
	Role      int	 `gorm:"not null"`
	Status    string `gorm:"not null;default:active"`
	LastLogin int64
}

type Role struct {
	gorm.Model
	ID       int    `gorm:"primaryKey"`
	RoleType int
	RoleName string `gorm:"uniqueindex;not null;size:255"`
	Status   string `gorm:"not null;default:active"`
}

// 应用Model
type App struct {
	gorm.Model
	// 同一类别下的排序位置
	Order	    int	   `gorm:"not null"`
	Name        string `gorm:"uniqueindex;not null;size:36"`
	Description string `gorm:"not null;size:255"`
	Icon        string `gorm:"not null;size:255"`
	// 路由模式
	Mode 		string `gorm:"not null;size:36"`
	Path        string `gorm:"not null;size:255"`
	Role		int    `gorm:"not null"`
	CategoryID  int    `gorm:"not null"`
}


// 应用分类
type AppCategory struct {
	gorm.Model
	Icon  string `gorm:"not null;size:255"`
	Name  string `gorm:"uniqueindex;not null;size:36"`
	// 排序顺序
	Order int    `gorm:"not null"`
}
