package model

import (
	"log"
	"our_blog/dao"
	"sync"
	"time"
)

type User struct {
	UserId     int       `gorm:"column:user_id"`
	Username   string    `gorm:"column:username"`
	Password   string    `gorm:"column:password"`
	Email      string    `gorm:"column:email"`
	CreateTime time.Time `gorm:"column:create_time"`
}

func (User) TableName() string {
	return "user"
}

type UserDao struct {
}

var userDao *UserDao
var userOnce sync.Once

func NewUserDaoInstance() *UserDao {
	userOnce.Do(func() {
		userDao = &UserDao{}
	})
	return userDao
}

func (UserDao) CreateUser(user *User) (err error) {
	err = dao.DB.Create(user).Error
	if err != nil {
		log.Println("create user failed, err : ", err)
		return err
	}
	return nil
}

func (UserDao) GetUserById(userId int) (user User, err error) {
	err = dao.DB.Where("user_id =?", userId).First(&user).Error
	if err != nil {
		log.Println("get user by id failed, err : ", err)
		return user, err
	}
	return user, nil
}

// 写一个重置密码的函数
func (UserDao) UpdateUserPassword(userId int, password string) (err error) {
	err = dao.DB.Model(&User{}).Where("user_id =?", userId).Update("password", password).Error
	if err != nil {
		log.Println("update user password failed, err : ", err)
		return err
	}
	return nil
}
