package dao

import (
	"gorm.io/gorm"
	"our_blog/db"
	"our_blog/model"
	"time"
)

type LikeDao struct{}

func NewLikeDaoInstance() *LikeDao {
	return &LikeDao{}
}

func (d *LikeDao) CreateLike(articleID int64, userID int64) error {
	like := model.Like{
		ArticleID: articleID,
		UserID:    userID,
		CreatedAt: time.Now(),
	}
	result := db.DB.Create(&like)
	return result.Error
}

func (d *LikeDao) DeleteLike(articleID int64, userID int64) error {
	result := db.DB.Where("article_id = ? AND user_id = ?", articleID, userID).Delete(&model.Like{})
	return result.Error
}

func (d *LikeDao) CheckLike(articleID int64, userID int64) (bool, error) {
	var like model.Like
	result := db.DB.Where("article_id = ? AND user_id = ?", articleID, userID).First(&like)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return false, nil
		}
		return false, result.Error
	}
	return true, nil
}
