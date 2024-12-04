package handler

import (
	"github.com/gin-gonic/gin"
	"log"
	"our_blog/commen/result"
	"our_blog/model/dto"
	"our_blog/service/article"
)

func ArticlePublishHandler(c *gin.Context) {
	var a dto.ArticlePublishReq
	err := c.ShouldBind(&a)
	if err != nil {
		log.Println("get article info failed, err:", err)
		result.Error(c, result.GetReqErrStatus)
		return
	}
	userid, ok := c.Get("userid")
	if !ok {
		log.Println("get userid failed")
		result.Error(c, result.ServerErrStatus)
		return
	}
	if userid, ok := userid.(int64); ok {
		data, err := article.ArticlePublish(a, userid)
		if err != nil {
			result.Error(c, result.ArticlePubErrStatus)
			return
		}
		r := make(result.R)
		r.ToMap(data)
		result.Sucess(c, r)
	}
}
