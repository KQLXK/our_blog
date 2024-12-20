package model

import (
	"gorm.io/gorm"
	"time"
)

// Like 定义点赞模型
type Like struct {
	gorm.Model
	DeletedAt gorm.DeletedAt `gorm:"index;comment:软删除时间"`
	ArticleID int64          `gorm:"index;comment:文章ID"`
	UserID    int64          `gorm:"index;comment:用户ID"`
	CreatedAt time.Time      `gorm:"comment:创建时间"`
	UpdatedAt time.Time      `gorm:"comment:更新时间"`
}
