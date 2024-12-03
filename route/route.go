package route

import (
	"github.com/gin-gonic/gin"
	"our_blog/handler"
	"our_blog/middleware"
)

func SetUpRouter() *gin.Engine {

	r := gin.Default()

	r.Use(middleware.Auth())

	userGroup := r.Group("/user")
	{
		//用户注册
		userGroup.POST("/register", handler.UserRegisterHandler)

		//用户登录
		userGroup.POST("/login", handler.UserLoginHandler)

	}

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"msg": "sucess",
		})
	})

	return r

}
