package handler

import (
	"github.com/gin-gonic/gin"
	"log"
	"our_blog/commen/result"
	"our_blog/model/dto"
	"our_blog/service/user"
)

func UserRegisterHandler(c *gin.Context) {
	var u dto.UserRegisterRequest
	err := c.ShouldBind(&u)
	if err != nil {
		log.Println("get userinfo failed, err: ", err)
		result.Error(c, result.GetReqErrStatus)
		return
	}

	data, err := user.UserRegister(&u)
	if err != nil {
		if err == user.UsernameExistErr {
			result.Error(c, result.UsernameExistErrStatus)
			return
		} else if err == user.EmailExistErr {
			result.Error(c, result.EmailExistErrStatus)
			return
		} else {
			result.Error(c, result.RegisterErrStatus)
			return
		}
	} else {
		result.Sucess(c, data)
		return
	}

}

func UserLoginHandler(c *gin.Context) {
	var u dto.UserLoginRequest
	err := c.ShouldBind(&u)
	if err != nil {
		log.Println("get userinfo failed, err: ", err)
		result.Error(c, result.GetReqErrStatus)
		return
	}

	data, err := user.UserLogin(u)
	if err != nil {
		if err == user.UsernameNotExistErr {
			result.Error(c, result.UsernameNotExsitsErrStatus)
		} else if err == user.PasswordWrongErr {
			result.Error(c, result.PasswordWrongErrStatus)
		} else {
			result.Error(c, result.ServerErrStatus)
		}
	} else {
		result.Sucess(c, data)
		return
	}
}
