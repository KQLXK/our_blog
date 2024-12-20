package dto

import (
	"gorm.io/gorm"
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
}

type Like struct {
	gorm.Model
	DeletedAt gorm.DeletedAt `gorm:"index;comment:软删除时间" json:"deleted_at"`
	ArticleID int64          `gorm:"index;comment:文章ID" json:"article_id"`
	UserID    int64          `gorm:"index;comment:用户ID" json:"user_id"`
	CreatedAt time.Time      `gorm:"comment:创建时间" json:"created_at"`
	UpdatedAt time.Time      `gorm:"comment:更新时间" json:"updated_at"`
}
