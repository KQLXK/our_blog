package route

import (
	"github.com/gin-gonic/gin"
	"our_blog/handler"
)

func SetUpRouter() *gin.Engine {

	r := gin.Default()

	userGroup := r.Group("/user")
	{

		//用户注册
		userGroup.POST("/register", handler.UserRegisterHandler)

	}

	return r

}
