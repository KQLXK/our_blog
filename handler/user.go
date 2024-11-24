package handler

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"our_blog/dto"
	"our_blog/service/user"
)

func UserRegisterHandler(c *gin.Context) {
	var u dto.UserRegisterRequest
	if err := c.ShouldBind(&u); err != nil {
		log.Println("get userinfo failed, err: ", err)
		c.JSON(http.StatusBadRequest, dto.UserRegisterResponse{
			Status:  "error",
			Message: err.Error(),
		})
		return
	}

	data, err := user.UserRegister(&u)
	if err != nil {
		c.JSON(http.StatusBadRequest, data)
		return
	} else {
		c.JSON(http.StatusOK, data)
	}

}
