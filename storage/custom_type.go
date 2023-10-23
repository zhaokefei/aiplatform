package storage

const (
	Guest       string = "guest"
	RegularUser string = "regular_user"
	Developer   string = "developer"
	SystemAdmin string = "system_admin"
	SuperAdmin  string = "super_admin"
)

var Roles = []string{
	Guest,
	RegularUser,
	Developer,
	SystemAdmin,
	SuperAdmin,
}
