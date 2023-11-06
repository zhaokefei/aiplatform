package storage


// 权限相关
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


// 应用分类相关
const (
	SystemID int = 1
	SystemName = "系统设置"
	SystemOrder = 999
	DefaultID int = 2
	DefaultName = "默认分类"
	DefaultOrder = 0
)


type AppMode string

const (
	InnerMode AppMode = "内部路由"
	OutterMode AppMode = "外部路由"
	AppCenterID int = 1
	RoleCenterID int = 2
)
