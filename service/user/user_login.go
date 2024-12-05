package user

import (
	"errors"
	"our_blog/commen/utils"
	"our_blog/model/dao"
	"our_blog/model/dto"
)

type UserLoginFlow struct {
	username string
	password string
}

//func UserLogin(ul dto.UserLoginRequest)

var (
	UsernameNotExistErr = errors.New("username exist already")
	PasswordWrongErr    = errors.New("password wrong")
)

func UserLogin(ul dto.UserLoginRequest) (*dto.UserLoginResponse, error) {
	return NewUserLoginFlow(ul).Do()
}

func NewUserLoginFlow(ul dto.UserLoginRequest) *UserLoginFlow {
	return &UserLoginFlow{
		username: ul.Username,
		password: ul.Password,
	}
}

func (f *UserLoginFlow) Do() (*dto.UserLoginResponse, error) {
	//检查用户名是否存在
	err := f.checkdata()
	if err != nil {
		return nil, err
	}

	//检查密码正确
	user, err := f.login()
	if err != nil {
		return nil, err
	}

	//创建token
	acctoken, err := utils.CreateUserToken(user.UserId)
	if err != nil {
		return nil, err
	}
	return &dto.UserLoginResponse{
		UserId:      user.UserId,
		Username:    user.Username,
		Email:       user.Email,
		CreateTime:  user.CreateTime,
		AccessToken: acctoken,
	}, nil

}

func (f *UserLoginFlow) checkdata() error {
	if _, err := dao.NewUserDaoInstance().GetUserByUsername(f.username); err != nil {
		return UsernameNotExistErr
	}
	return nil
}

func (f *UserLoginFlow) login() (*dao.User, error) {
	user, _ := dao.NewUserDaoInstance().GetUserByUsername(f.username)
	if f.password == user.Password {
		return &user, nil
	}
	return nil, PasswordWrongErr
}
