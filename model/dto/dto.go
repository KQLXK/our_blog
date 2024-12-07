package dto

import "time"

type UserRegisterRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type UserRegisterResponse struct {
	UserId      int64     `json:"userId"`
	Username    string    `json:"username"`
	Email       string    `json:"email"`
	CreateTime  time.Time `json:"create_Time"`
	AccessToken string    `json:"access_token"`
}

type UserLoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserLoginResponse struct {
	UserId      int64     `json:"userId"`
	Username    string    `json:"username"`
	Email       string    `json:"email"`
	CreateTime  time.Time `json:"create_Time"`
	AccessToken string    `json:"access_token"`
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
	UpdateAt  time.Time `json:"Update_at"`
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
}
