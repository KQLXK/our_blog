package handler

import (
	"github.com/gin-gonic/gin"
	"log"
	"our_blog/commen/result"
	"our_blog/dto"
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
		} else if err == user.EmailExistErr {
			result.Error(c, result.EmailExistErrStatus)
		} else {
			result.Error(c, result.RegisterErrStatus)
		}
	} else {
		r := make(result.R)
		r.ToMap(data)
		result.Sucess(c, r)
	}

}

func UserLoginHandler(c *gin.Context) {
	var u dto.UserLoginRequest
	err := c.ShouldBind(&u)
	if err != nil {
		log.Println("get userinfo failed, err: ", err)
		result.Error(c, result.GetReqErrStatus)
	}

	data, err := user.NewUserLoginFlow(u).Do()
	if err != nil {
		if err == user.UsernameNotExistErr {
			result.Error(c, result.UsernameNotExsitsErrStatus)
		} else if err == user.PasswordWrongErr {
			result.Error(c, result.PasswordWrongErr)
		} else {
			result.Error(c, result.ServerErrStatus)
		}
	} else {
		r := make(result.R)
		r.ToMap(data)
		result.Sucess(c, r)
	}
}
