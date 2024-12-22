package handler

import (
	"github.com/gin-gonic/gin"
	"log"
	"our_blog/commen/result"
	"our_blog/service/comment"
	"strconv"
)

func CommentQueryHandler(c *gin.Context) {
	commentIdStr := c.Param("comment_id") // 从 URL 获取评论 ID

	// 将字符串转换为 int64
	commentId, err := strconv.ParseInt(commentIdStr, 10, 64)
	if err != nil {
		log.Println("invalid comment id, err:", err)
		result.Error(c, result.GetReqErrStatus)
		return
	}

	data, err := comment.CommentQuery(commentId)
	if err != nil {
		log.Println("get comment failed, err:", err)
		result.Error(c, result.ServerErrStatus)
		return
	}

	result.Sucess(c, data) // 返回成功的响应
}

func CommentListHandler(c *gin.Context) {
	ArticleIdStr := c.Param("article_id")
	ArticleId, err := strconv.ParseInt(ArticleIdStr, 10, 64)
	if err != nil {
		result.Error(c, result.ServerErrStatus)
		return
	}
	data, err := comment.CommentList(ArticleId)
	if err != nil {
		result.Error(c, result.ServerErrStatus)
		return
	}
	result.Sucess(c, data)
}
