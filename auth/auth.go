package auth


type Auth interface {
	// 登陆
	Login(string) (string, error)
	// 登出
	Logout() (bool, error)
	// 是否登录
	IsLogin() (bool, error)
}