package storage

import "gorm.io/gorm"


type User struct {
	gorm.Model
	Username  string `gorm:"uniqueindex;not null;size:255"`
	Password  string
	Age       string
	Sex       string
	Email     string `gorm:"uniqueindex;not null;size:255"`
	Role      int	 
	Status    string `gorm:"not null;default:active"`
	LastLogin int64
}

type Role struct {
	gorm.Model
	ID       int      `gorm:"primaryKey"`
	RoleName UserRole `gorm:"uniqueindex;not null;size:255"`
	Status   string   `gorm:"not null;default:active"`
}
