package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"log"
	"our_blog/commen/result"
	"our_blog/model/dto"
	"our_blog/service/article"
	"strconv"
)

func ArticlePublishHandler(c *gin.Context) {
	var a dto.ArticlePublishReq
	if err := c.ShouldBind(&a); err != nil {
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
	//类型断言
	userId, ok := userid.(int64)
	if !ok {
		log.Println("userid format err")
		result.Error(c, result.ServerErrStatus)
		return
	}
	data, err := article.ArticlePublish(a, userId)
	if err != nil {
		result.Error(c, result.ArticlePubErrStatus)
		return
	}
	result.Sucess(c, data)
	return
}

func ArticleUpdateHandler(c *gin.Context) {
	var a dto.ArticleUpdateReq
	if err := c.ShouldBind(&a); err != nil {
		log.Println("get article update info failed, err:", err)
		result.Error(c, result.GetReqErrStatus)
		return
	}
	log.Println(a)
	userid, ok := c.Get("userid")
	if !ok {
		log.Println("get userid failed")
		result.Error(c, result.ServerErrStatus)
		return
	}
	//类型断言
	userId, ok := userid.(int64)
	if !ok {
		log.Println("userid format err")
		result.Error(c, result.ServerErrStatus)
		return
	}
	data, err := article.ArticleUpdate(a, userId)
	if err != nil {
		if err == article.UpdateUnautherizedErr {
			result.Error(c, result.UnauthorizedStatus)
			return
		} else if err == article.ArticleNotFoundErr {
			result.Error(c, result.ArticleNotFoundErrStatus)
			return
		} else {
			result.Error(c, result.ServerErrStatus)
			return
		}
	}
	result.Sucess(c, data)
	return
}

func ArtQueryByIdHandler(c *gin.Context) {
	ArticleId := c.Param("articleid")
	if ArticleId == "" {
		log.Println("get article_id failed")
		result.Error(c, result.GetReqErrStatus)
		return
	}
	articleid, err := strconv.ParseInt(ArticleId, 10, 64)
	log.Println("articleid:", ArticleId)
	if err != nil {
		result.Error(c, result.ServerErrStatus)
		return
	}
	data, err := article.ArtQueryById(articleid)
	if err != nil {
		if err == article.ArticleNotFoundErr {
			result.Error(c, result.ArticleNotFoundErrStatus)
			return
		} else if err == article.BadDataErr {
			result.Error(c, result.InvalidDataErrStatus)
		} else {
			result.Error(c, result.ServerErrStatus)
			return
		}
	}
	result.Sucess(c, data)
	return
}

func ArtQueryByPageHandler(c *gin.Context) {
	// 获取查询参数并转换为 int64
	pageStr := c.Query("page")
	Page, _ := strconv.ParseInt(pageStr, 10, 64)

	pageSizeStr := c.Query("pagesize")
	PageSize, _ := strconv.ParseInt(pageSizeStr, 10, 64)

	UserId, err := GetUserid(c)
	if err != nil {
		return
	}
	// 调用 Article 查询方法
	articles, err := article.QueryByPage(int(Page), int(PageSize), UserId)
	if err != nil {
		if err == article.BadDataErr {
			result.Error(c, result.InvalidDataErrStatus)
			return
		} else {
			result.Error(c, result.ServerErrStatus)
			return
		}
	}
	result.Sucess(c, articles)
}

func DeleteArticleHandler(c *gin.Context) {
	UserId, err := GetUserid(c)
	if err != nil {
		result.Error(c, result.ServerErrStatus)
		return
	}
	articleid := c.Query("article_id")
	ArticleId, _ := strconv.ParseInt(articleid, 10, 64)
	err = article.ArticleDelete(ArticleId, UserId)
	if err != nil {
		if err == article.ArticleNotFoundErr {
			result.Error(c, result.ArticleNotFoundErrStatus)
			return
		} else {
			result.Error(c, result.ServerErrStatus)
			return
		}
	}
	result.Sucess(c, nil)
}

func GetUserid(c *gin.Context) (int64, error) {
	userid, ok := c.Get("userid")
	if !ok {
		log.Println("get userid failed")
		result.Error(c, result.ServerErrStatus)
		return -1, errors.New("get userid failed")
	}
	UserId, ok := userid.(int64)
	if !ok {
		log.Println("userid format err")
		result.Error(c, result.ServerErrStatus)
		return -1, errors.New("userid format err")
	}
	return UserId, nil
}
