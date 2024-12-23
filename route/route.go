package route

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"our_blog/handler"
	"our_blog/middleware"
)

func SetUpRouter() *gin.Engine {

	r := gin.Default()

	// 配置 CORS 中间件
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://127.0.0.1:5500"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"content-type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	r.Use(middleware.Auth())

	userGroup := r.Group("/user")
	{
		//用户注册
		userGroup.POST("/register", handler.UserRegisterHandler)
		//用户登录
		userGroup.POST("/login", handler.UserLoginHandler)

		//重置密码
		userGroup.POST("/reset", handler.UserResetPasswordHandler)
		//使用id获取用户信息
		userGroup.GET("/:id", handler.UserGetByIdHandler)
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
		//点赞
		ArticleGroup.POST("/:article_id/like", handler.ArticleLikeHandler)
		//文章评论
		ArticleGroup.POST("/:article_id/comment", handler.ArticleCommentHandler)
		//搜索文章
		ArticleGroup.GET("/search", handler.ArticleSearchHandler)
		//文章分类
		ArticleGroup.GET("/category", handler.ArticleCategoryHandler)
	}

	CommentGroup := r.Group("/comment")
	{
		//按文章id获取comment
		CommentGroup.GET("/list/:article_id", handler.CommentListHandler)
		//查询评论
		CommentGroup.GET("/query/:comment_id", handler.CommentQueryHandler)
	}

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"msg": "sucess",
		})
	})

	return r

}
