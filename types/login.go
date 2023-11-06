package types

type LoginBody struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type RegisterBody struct {
	Username      string `json:"username"`
	Password      string `json:"password"`
	AgainPassword string `json:"again_password"`
	Email         string `json:"email"`
	Sex           string `json:"sex"`
	Age           string `json:"age"`
}


type UserRoleBody struct {
	Username string `json:"username"`
	Rolename string `json:"rolename"`
}
