package dto

import (
	"our_blog/model/dao"
	"time"
)

type UserRegisterRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type UserRegisterResponse struct {
	UserId      int64     `json:"userid"`
	Username    string    `json:"username"`
	Email       string    `json:"email"`
	CreateTime  time.Time `json:"create_time"`
	AccessToken string    `json:"access_token"`
}

type UserLoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserLoginResponse struct {
	UserId      int64     `json:"userid"`
	Username    string    `json:"username"`
	Email       string    `json:"email"`
	CreateTime  time.Time `json:"create_time"`
	AccessToken string    `json:"access_token"`
}

type UserResetPasswordRequest struct {
	Username    string `json:"username"`
	NewPassword string `json:"new_password"`
}

type ArticlePublishReq struct {
	Title    string `json:"title"`    // 文章标题
	Username string `json:"username"` // 文章作者
	Excerpt  string `json:"excerpt"`  // 文章摘要
	Category string `json:"category"` // 文章分类
	Content  string `json:"content"`  // 文章内容
	Status   string `json:"status"`   // 文章状态
}

type ArticlePublishResp struct {
	ArticleId int64     `json:"article_id"`
	Title     string    `json:"title"`
	CreateAt  time.Time `json:"create_at"`
}

type ArticleUpdateReq struct {
	ArticleId int64  `json:"article_id"`
	Title     string `json:"title"`    // 文章标题
	Excerpt   string `json:"excerpt"`  // 文章摘要
	Category  string `json:"category"` // 文章分类
	Content   string `json:"content"`  // 文章内容
	Status    string `json:"status"`   // 文章状态
}

type ArticleUpdateResp struct {
	ArticleId int64     `json:"article_id"`
	Title     string    `json:"title"`
	UpdateAt  time.Time `json:"update_at"`
}

type ArtQueryByIdResp struct {
	Title      string    `json:"title"`
	ArticleId  int64     `json:"article_id"`
	UserId     int64     `json:"user_id"`
	Excerpt    string    `json:"excerpt"`
	Category   string    `json:"category"`
	Content    string    `json:"content"`
	Status     string    `json:"status"`
	CreateTime time.Time `json:"create_time"`
	UpdateTime time.Time `json:"update_time"`
	LikeCount  int       `json:"like_count"`
	View       int64     `json:"view"`
}

type ArticleCommentReq struct {
	ArticleId int64  `json:"article_id"`
	Content   string `json:"content"`
	ParentId  *int64 `json:"parent_id"`
}

type ArticleCommentResp struct {
	CommentId  int64     `json:"comment_id"`
	CreateTime time.Time `json:"create_time"`
}

type CommentGetResp struct {
	CommentId  int64     `json:"comment_id"`
	UserId     int64     `json:"user_id"`
	ArticleId  int64     `json:"article_id"`
	ParentId   *int64    `json:"parent_id"`
	Content    string    `json:"content"`
	CreateTime time.Time `json:"create_time"`
}

//// CommentWithReplies 包含评论及其回复的结构
//type CommentWithReplies struct {
//	Comment *dao.Comment   `json:"comment"`
//	Replies []*dao.Comment `json:"replies"` // 每个评论的回复
//}

// CommentListResp 返回评论列表的结构
type CommentListResp struct {
	Comments []*dao.Comment `json:"comment"`
}
