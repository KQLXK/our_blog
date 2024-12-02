package user

import (
	"errors"
	"our_blog/dto"
	"our_blog/repository"
	"time"
)

var (
	UsernameExistErr = errors.New("username exist already")
	EmailExistErr    = errors.New("email was registered")
)

type UserRegisterFlow struct {
	userid      int64
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
		return nil, err
	}
	if err := f.register(); err != nil {
		return nil, err
	}
	return &dto.UserRegisterResponse{
		UserId:     f.userid,
		Username:   f.username,
		Email:      f.email,
		CreateTime: f.create_time,
		//"token": "eyJhbGciOiJIUzI1NiIsInR..."
	}, nil
}

func (f *UserRegisterFlow) checkData() error {
	if _, err := repository.NewUserDaoInstance().GetUserByUsername(f.username); err == nil {
		return UsernameExistErr
	}
	if _, err := repository.NewUserDaoInstance().GetUserByEmail(f.email); err == nil {
		return EmailExistErr
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
