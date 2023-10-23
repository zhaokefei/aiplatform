package storage

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

var ctx = context.Background()

type UserOps struct {
	Username string
	AuthKey  string
	UserData *User
}

// New creates a new UserOps instance with the given username.
//
// It checks if the user exists and updates the UserData.
//
// Parameters:
// - username: The username of the user.
//
// Returns:
// - *UserOps: The UserOps instance.
// - error: An error if the user doesn't exist or if there was an error retrieving the user information.
func NewUserOps(username string) (*UserOps, error) {
	ops := &UserOps{
		Username: username,
		AuthKey:  "auth:login:" + username,
	}
	// 检查User 是否存在
	user, err := ops.UserInfo()
	if err != nil || user == nil {
		return nil, err
	}
	// 更新UserData
	ops.UserData = user
	return ops, nil
}

// Logined returns the token associated with a logged-in user and any error encountered.
//
// It takes no parameters.
// It returns a string, the token associated with the logged-in user, and an error, if any.
func (uo *UserOps) Logined() (string, error) {
	val, err := RedisClient.Get(ctx, uo.AuthKey).Result()
	if err != nil || err == redis.Nil {
		return "", nil
	}
	return val, nil
}

// Logout logs out the user.
//
// This function does not take any parameters.
// It returns a boolean value indicating whether the logout was successful or not, and an error if there was any.
func (uo *UserOps) Logout() (bool, error) {
	err := RedisClient.Del(ctx, uo.AuthKey).Err()
	if err != nil {
		return false, err
	}
	return true, nil
}

// Login verifies the given password and generates a session ID for the user.
//
// Parameters:
// - password: The password to be verified.
//
// Returns:
// - string: The generated session ID.
// - error: An error if the password is incorrect or there is an issue with the Redis client.
func (uo *UserOps) Login(password string) (string, error) {
	if password != uo.UserData.Password {
		return "", errors.New("password is wrong")
	}
	SessionID := GenSessionId()
	err := RedisClient.Set(ctx, uo.AuthKey, SessionID, 24*3600).Err()
	if err != nil {
		return "", err
	}
	return SessionID, nil
}

// UserInfo retrieves the user information.
//
// It does not take any parameters.
// It returns a pointer to a User struct and an error.
func (uo *UserOps) UserInfo() (*User, error) {
	var user User
	if uo.UserData != nil {
		return uo.UserData, nil
	}
	result := MysqlClient.Where("Username = ?", uo.Username).First(&user)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, result.Error
		}
		return nil, result.Error
	}
	return &user, nil
}

// SetRole sets the role of a user.
//
// name: the name of the role to be set.
// Returns a boolean indicating whether the role was set successfully and an error if any.
func (uo *UserOps) SetRole(name string) (bool, error) {
	var role Role
	// 判断是否是符合条件的role类型
	if !IsValidRole(name) {
		return false, errors.New("invalid role")
	}
	result := MysqlClient.Where(&Role{RoleName: name}).First(&role)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return false, result.Error
		}
		return false, result.Error
	}
	uo.UserData.Role = int(role.ID)
	return true, nil
}

// UserRegister registers a new user with the provided username, password, email, and additional parameters.
//
// The function takes the following parameters:
// - username: a string representing the username of the user.
// - password: a string representing the password of the user.
// - email: a string representing the email address of the user.
// - params: a map[string]string containing additional parameters for the user.
//
// The function returns a boolean value indicating whether the user registration was successful, and an error value
// in case of any errors.
func UserRegister(username, password, email string, params map[string]string) (bool, error) {
	// 内置函数，获取默认值
	MapDefault := func(key string, defaultValue string) string {
		value, ok := params[key]
		if !ok {
			return defaultValue
		}
		return value
	}
	role, err := RoleInfo(RegularUser)
	if err != nil {
		return false, errors.New("未找到合适的用户角色")
	}
	// 构造user信息
	user := User{
		Username: username,
		Password: password,
		Email:    email,
		Age:      MapDefault("age", ""),
		Sex:      MapDefault("sex", ""),
		Role:     role.ID,
	}
	// 创建User
	result := MysqlClient.Create(&user)
	err = result.Error
	if err != nil {
		return false, err
	}
	return true, result.Error
}

func GenSessionId() string {
	return uuid.New().String()
}
