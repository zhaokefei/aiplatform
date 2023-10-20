package auth

import (
	"errors"

	DB "github.com/zhaokefei/aiplatform/storage"
)

type User struct {
	Username string
	UserInfo *DB.UserOps
}

func NewUser(username string) (*User, error) {
	uo, err := DB.NewUserOps(username)
	if err != nil && uo == nil {
		return nil, errors.New("username doesn't exist")
	}
	return &User{
		Username: username,
		UserInfo: uo,
	}, nil
}

// Login is a function that authenticates a user with a username and password and returns a token and an error.
//
// Parameters:
// - username: a string representing the username of the user.
// - password: a string representing the password of the user.
//
// Returns:
// - token: a string representing the authentication token.
// - err: an error indicating any authentication errors.
func (u *User) Login(password string) (token string, err error) {
	if u.Username == "" || password == "" {
		return "", errors.New("username or password is empty")
	}
	token, err = u.UserInfo.Logined()
	if err != nil || token == "" {
		return u.UserInfo.Login(password)
	}
	return token, nil
}

// Logout logs out a user.
//
// It takes a username as a parameter and returns the logout status
// as a boolean and any error encountered during the process.
func (u *User) Logout() (status bool, err error) {
	if u.Username == "" {
		return false, errors.New("username is empty")
	}
	return u.UserInfo.Logout()
}

// IsLogin checks if the user with the given username is logged in.
//
// Parameter:
// - username: the username of the user to check.
//
// Return:
// - status: true if the user is logged in, false otherwise.
func (u *User) IsLogin() (status bool, err error) {
	if u.Username == "" {
		return false, errors.New("username is empty")
	}
	token, err := u.UserInfo.Logined()
	if err != nil || token == "" {
		return false, nil
	}
	return true, nil
}

// Register is a function that registers a user with a username, password, and additional parameters.
//
// It takes the following parameters:
// - username: a string representing the username of the user to be registered.
// - password: a string representing the password of the user to be registered.
// - again_password: a string representing the password confirmation.
// - params: a map[string]string representing any additional parameters for the user registration.
//
// It returns:
// - status: a boolean indicating the success of the registration.
// - err: an error indicating any registration errors that occurred.
func Register(username, password, again_password, email string, params map[string]string) (status bool, err error) {
	if username == "" || password == "" || again_password == "" {
		return false, errors.New("username or password is empty")
	} else if password != again_password {
		return false, errors.New("password not match")
	}
	return DB.UserRegister(username, password, email, params)
}
