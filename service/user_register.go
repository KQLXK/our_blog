package service

import (
	"errors"
	"our_blog/dto"
	"our_blog/repository"
	"time"
)

type UserRegisterFlow struct {
	userid      int
	username    string
	password    string
	email       string
	create_time time.Time
}

func UserRegister(ur *dto.UserRegisterRequest) (*dto.UserRegisterResponse, error) {
	return NewUserRegisterFlow(ur).Do()
}

func NewUserRegisterFlow(ur *dto.UserRegisterRequest) *UserRegisterFlow {
	return &UserRegisterFlow{
		username:    ur.Username,
		password:    ur.Password,
		email:       ur.Email,
		create_time: time.Now(),
	}
}

func (f *UserRegisterFlow) Do() (*dto.UserRegisterResponse, error) {
	if err := f.checkData(); err != nil {
		return &dto.UserRegisterResponse{
			Status:  "error",
			Message: err.Error(),
		}, err
	}
	if err := f.register(); err != nil {
		return &dto.UserRegisterResponse{
			Status:  "error",
			Message: err.Error(),
		}, err
	}
	return &dto.UserRegisterResponse{
		Status:  "sucess",
		Message: "注册成功",
		Data: map[string]interface{}{
			"userId":     f.userid,
			"username":   f.username,
			"email":      f.email,
			"createTime": f.create_time,
			//"token": "eyJhbGciOiJIUzI1NiIsInR..."
		},
	}, nil
}

func (f *UserRegisterFlow) checkData() error {
	if _, err := repository.NewUserDaoInstance().GetUserByUsername(f.username); err == nil {
		return errors.New("username exist already")
	}
	if _, err := repository.NewUserDaoInstance().GetUserByEmail(f.email); err == nil {
		return errors.New("email was registered")
	}
	return nil
}

func (f *UserRegisterFlow) register() error {
	user := &repository.User{
		Username:   f.username,
		Password:   f.password,
		Email:      f.email,
		CreateTime: f.create_time,
	}
	if err := repository.NewUserDaoInstance().CreateUser(user); err != nil {
		return err
	}
	f.userid = user.UserId
	return nil
}
