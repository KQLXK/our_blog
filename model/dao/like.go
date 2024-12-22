package dao

import (
	"gorm.io/gorm"
	"our_blog/db"
	"our_blog/model/dto"
	"time"
)

type Like struct {
	gorm.Model
	DeletedAt gorm.DeletedAt `gorm:"index;comment:软删除时间" json:"deleted_at"`
	ArticleID int64          `gorm:"index;comment:文章ID" json:"article_id"`
	UserID    int64          `gorm:"index;comment:用户ID" json:"user_id"`
	CreatedAt time.Time      `gorm:"comment:创建时间" json:"created_at"`
	UpdatedAt time.Time      `gorm:"comment:更新时间" json:"updated_at"`
}

type LikeDao struct{}

func NewLikeDaoInstance() *LikeDao {
	return &LikeDao{}
}

func (d *LikeDao) CreateLike(articleID int64, userID int64) error {
	like := dto.Like{
		ArticleID: articleID,
		UserID:    userID,
		CreatedAt: time.Now(),
	}
	result := db.DB.Create(&like)
	return result.Error
}

func (d *LikeDao) DeleteLike(articleID int64, userID int64) error {
	result := db.DB.Where("article_id = ? AND user_id = ?", articleID, userID).Delete(&dto.Like{})
	return result.Error
}

func (d *LikeDao) CheckLike(articleID int64, userID int64) (bool, error) {
	var like dto.Like
	result := db.DB.Where("article_id = ? AND user_id = ?", articleID, userID).First(&like)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return false, nil
		}
		return false, result.Error
	}
	return true, nil
}
