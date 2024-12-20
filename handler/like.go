package handler

import (
	"github.com/gin-gonic/gin"
	"log"
	"our_blog/commen/result"
	"our_blog/model/dao"
	"strconv"
)

// ArticleLikeHandler 处理文章点赞
func ArticleLikeHandler(c *gin.Context) {
	articleIDStr := c.Param("article_id")
	userIDStr, _ := GetUserid(c) //户ID通过请求头传递

	articleID, err := strconv.ParseInt(articleIDStr, 10, 64)
	if err != nil {
		log.Println("invalid article id, err:", err)
		result.Error(c, result.InvalidDataErrStatus)
		return
	}

	userID, err := strconv.ParseInt(strconv.FormatInt(userIDStr, 10), 10, 64)
	if err != nil {
		log.Println("invalid user id, err:", err)
		result.Error(c, result.InvalidDataErrStatus)
		return
	}

	likeDao := dao.NewLikeDaoInstance()
	liked, err := likeDao.CheckLike(articleID, userID)
	if err != nil {
		result.Error(c, result.ServerErrStatus)
		return
	}

	if liked {
		result.Error(c, result.AlreadyLikedErr) //一个新的状态码表示已经点赞
		return
	}

	err = likeDao.CreateLike(articleID, userID)
	if err != nil {
		result.Error(c, result.ServerErrStatus)
		return
	}

	result.Sucess(c, gin.H{"message": "article liked"})
}
