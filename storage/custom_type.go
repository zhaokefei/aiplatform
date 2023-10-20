package storage

import "database/sql/driver"

type UserRole string

const (
	Guest       UserRole = "guest"
	RegularUser UserRole = "regular_user"
	Developer   UserRole = "developer"
	SystemAdmin UserRole = "system_admin"
	SuperAdmin  UserRole = "super_admin"
)

func (r *UserRole) Scan(value interface{}) error {
	*r = UserRole(value.(string))
	return nil
}

func (r UserRole) Value() (driver.Value, error) {
	return string(r), nil
}
