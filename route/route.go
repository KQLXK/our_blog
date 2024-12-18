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

		//重置密码
		userGroup.POST("/reset", handler.UserResetPasswordHandler)
	}

	ArticleGroup := r.Group("/article")
	{
		//发表文章
		ArticleGroup.POST("/publish", handler.ArticlePublishHandler)
		//更新文章
		ArticleGroup.PUT("/update", handler.ArticleUpdateHandler)
		//获取文章-按页
		ArticleGroup.GET("/querybypage", handler.ArtQueryByPageHandler)
		//获取文章-按文章id
		ArticleGroup.GET("/querybyid/:articleid", handler.ArtQueryByIdHandler)
		//删除文章
		ArticleGroup.DELETE("/delete", handler.DeleteArticleHandler)
	}

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"msg": "sucess",
		})
	})

	return r

}
