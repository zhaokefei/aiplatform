package storage


const (
	_ = iota
	GuestID
	_ = iota
	RegularUserID
	_ = iota
	DeveloperID
	_ = iota
	SystemAdminID
	_ = iota
	SuperAdminID
)


const (
	Guest       string = "guest"
	RegularUser string = "regular_user"
	Developer   string = "developer"
	SystemAdmin string = "system_admin"
	SuperAdmin  string = "super_admin"
)


var RoleMapper = map[string]int{
	Guest:       GuestID,
	RegularUser: RegularUserID,
	Developer:   DeveloperID,
	SystemAdmin: SystemAdminID,
	SuperAdmin:  SuperAdminID,
}


var Roles = []string{
	Guest,
	RegularUser,
	Developer,
	SystemAdmin,
	SuperAdmin,
}
