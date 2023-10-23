package storage

import (
	"errors"

	"gorm.io/gorm"
)

func NewRole(name string) (*Role, error) {
	if !IsValidRole(name) {
		return nil, errors.New("invalid role")
	}
	role := Role{
		RoleType: RoleMapper[name],
		RoleName: name,
		Status:   "active",
	}
	result := MysqlClient.Create(&role)
	if result.Error != nil {
		return nil, result.Error
	}
	return &role, nil
}

func IsSet(name string) bool {
	var role Role
	result := MysqlClient.Where(&Role{RoleName: name}).First(&role)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return false
		}
		return false
	}
	return true
}

func IsValidRole(name string) bool {
	for _, role := range Roles {
		if role == name {
			return true
		}
	}
	return false
}


func RoleInfo(name string) (*Role, error) {
	var role Role
	result := MysqlClient.Where(&Role{RoleName: name}).First(&role)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, result.Error
		}
		return nil, result.Error
	}
	return &role, nil
}
