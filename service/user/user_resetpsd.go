package user

import (
	"errors"
	"log"
	"our_blog/model/dao"
)

//	func UserResetPassword(username string, newPassword string) error {
//		user, err := dao.NewUserDaoInstance().GetUserByUsername(username)
//		if err != nil {
//			return UsernameNotExistErr
//		}
//		user.Password = newPassword
//		err = dao.NewUserDaoInstance().UpdateUser(user)
//		if err != nil {
//			log.Println("err", err)
//			return err
//		}
//		log.Println("success")
//		return nil
//	}
var (
	PasswordSameErr = errors.New("password same")
)

type PasswordResetFlow struct {
	username    string
	newPassword string
}

func NewPasswordResetFlow(username string, newPassword string) *PasswordResetFlow {
	return &PasswordResetFlow{
		username:    username,
		newPassword: newPassword,
	}
}

func (f *PasswordResetFlow) Do() error {
	if err := f.checkData(); err != nil {
		return err
	}
	if err := f.checkNewPassword(); err != nil {
		return err
	}
	if err := f.resetPassword(); err != nil {
		return err
	}
	return nil
}

func (f *PasswordResetFlow) checkData() error {
	_, err := dao.NewUserDaoInstance().GetUserByUsername(f.username)
	if err != nil {
		return UsernameNotExistErr
	}
	return nil
}

func (f *PasswordResetFlow) checkNewPassword() error {
	same, err := dao.NewUserDaoInstance().CheckPassword(f.username, f.newPassword)
	if err != nil {
		return err
	}
	if same {
		return PasswordSameErr // 你需要定义这个错误状态
	}
	return nil
}

func (f *PasswordResetFlow) resetPassword() error {
	user, err := dao.NewUserDaoInstance().GetUserByUsername(f.username)
	if err != nil {
		log.Println("err", err)
		return err
	}
	user.Password = f.newPassword
	err = dao.NewUserDaoInstance().UpdateUserPassword(int(user.UserId), user.Password) // 修改这里以使用新的UpdateUserPassword方法
	if err != nil {
		log.Println("err", err)
		return err
	}
	log.Println("success")
	return nil
}

func UserResetPassword(username string, newPassword string) error {
	return NewPasswordResetFlow(username, newPassword).Do()
}
