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
	if err := c.ShouldBind(&u); err != nil {
		log.Println("get userinfo failed, err: ", err)
		result.Error(c, result.RegisterErrStatus)
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
