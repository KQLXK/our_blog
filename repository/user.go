package repository

import (
	"log"
	"sync"
	"time"
)

type User struct {
	UserId     int       `gorm:"column:user_id;primaryKey;autoIncrement"`
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
	err = DB.Create(user).Error
	if err != nil {
		log.Println("create user failed, err : ", err)
		return err
	}
	log.Printf("User created successfully: %+v\n", user)
	return nil
}

func (UserDao) GetUserByUsername(username string) (user User, err error) {
	err = DB.Where("username = ?", username).First(&user).Error
	if err != nil {
		log.Println("get user by username failed, err: ", err)
		return user, err
	}
	return user, nil
}

func (UserDao) GetUserById(userId int) (user User, err error) {
	err = DB.Where("user_id =?", userId).First(&user).Error
	if err != nil {
		log.Println("get user by id failed, err : ", err)
		return user, err
	}
	return user, nil
}

func (UserDao) GetUserByEmail(email string) (user User, err error) {
	err = DB.Where("email = ?", email).First(&user).Error
	if err != nil {
		log.Println("get user by email failed, err: ", err)
		return user, err
	}
	return user, nil
}

// 写一个重置密码的函数
func (UserDao) UpdateUserPassword(userId int, password string) (err error) {
	err = DB.Model(&User{}).Where("user_id =?", userId).Update("password", password).Error
	if err != nil {
		log.Println("update user password failed, err : ", err)
		return err
	}
	return nil
}
