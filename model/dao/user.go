package dao

import (
	"log"
	"our_blog/db"
	"sync"
	"time"
)

type User struct {
	UserId     int64     `gorm:"column:user_id;primaryKey;autoIncrement"`
	Username   string    `gorm:"column:username"`
	Password   string    `gorm:"column:password"`
	Email      string    `gorm:"column:email"`
	CreateTime time.Time `gorm:"column:create_time"`
	IsAdmin    bool      `gorm:"column:is_admin;default:false"`
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
	err = db.DB.Create(user).Error
	if err != nil {
		log.Println("create user failed, err : ", err)
		return err
	}
	log.Printf("User created successfully: %+v\n", user)
	return nil
}

func (UserDao) GetUserByUsername(username string) (user User, err error) {
	err = db.DB.Where("username = ?", username).First(&user).Error
	if err != nil {
		log.Println("get user by username failed, err: ", err)
		return user, err
	}
	log.Println("get user by username success")
	return user, nil
}

func (UserDao) GetUserById(userId int64) (user User, err error) {
	err = db.DB.Where("user_id =?", userId).First(&user).Error
	if err != nil {
		log.Println("get user by id failed, err : ", err)
		return user, err
	}
	log.Println("get user by id sucess")
	return user, nil
}

func (UserDao) GetUserByEmail(email string) (user User, err error) {
	err = db.DB.Where("email = ?", email).First(&user).Error
	if err != nil {
		log.Println("get user by email failed, err: ", err)
		return user, err
	}
	log.Println("get user by email sucess")
	return user, nil
}

// 写一个重置密码的函数
func (UserDao) UpdateUserPassword(userId int, password string) (err error) {
	err = db.DB.Model(&User{}).Where("user_id = ?", userId).Update("password", password).Error
	if err != nil {
		log.Println("update user password failed, err : ", err)
		return err
	}
	return nil
}

func (UserDao) UpdateUser(user User) error {
	err := db.DB.Save(user).Error
	if err != nil {
		log.Println("update user failed, err : ", err)
		return err
	}
	log.Printf("User updated successfully: %+v\n", user)
	return nil
}

// 检测新密码与原密码是否相同
func (UserDao) CheckPassword(username string, password string) (bool, error) {
	user, err := NewUserDaoInstance().GetUserByUsername(username)
	if err != nil {
		return false, err
	}
	return user.Password == password, nil
}

func (UserDao) IsAdmin(UserID int64) (bool, error) {
	user, err := NewUserDaoInstance().GetUserById(UserID)
	if err != nil {
		return false, err
	}
	return user.IsAdmin, nil
}
