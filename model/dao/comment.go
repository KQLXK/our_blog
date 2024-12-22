package dao

import (
	"log"
	"our_blog/db"
	"sync"
	"time"
)

type Comment struct {
	CommentId  int64      `gorm:"column:comment_id;primaryKey;autoIncrement"` // 评论 ID
	ArticleId  int64      `gorm:"column:article_id"`                          // 关联的文章 ID
	UserId     int64      `gorm:"column:user_id"`                             // 评论者 ID
	Content    string     `gorm:"column:content"`                             // 评论内容
	ParentId   *int64     `gorm:"column:parent_id"`                           // 父评论 ID，支持回复
	CreateTime time.Time  `gorm:"column:create_time"`                         // 创建时间
	Replies    []*Comment `gorm:"foreignkey:ParentId"`                        // 关联的子回复
	//Status     string    `gorm:"column:status"`                              // 评论状态（如待审核、已通过等）
}

func (Comment) Tablename() string {
	return "comment"
}

type CommentDao struct {
}

var (
	commentDao  *CommentDao
	commentOnce sync.Once
)

func NewCommentDaoInstance() *CommentDao {
	commentOnce.Do(func() {
		commentDao = &CommentDao{}
	})
	return commentDao
}

func (CommentDao) Create(comment *Comment) (err error) {
	if err = db.DB.Create(comment).Error; err != nil {
		log.Println("create comment failed, err:", err)
		return err
	}
	log.Println("create comment success")
	return nil
}

func (CommentDao) Delete(commentId int64) (err error) {
	if err = db.DB.Model(&Comment{}).Delete(&Comment{}, commentId).Error; err != nil {
		log.Println("delete comment failed, err:", err)
		return err
	}
	log.Println("delete comment success")
	return nil
}

func (CommentDao) QueryById(commentId int64) (*Comment, error) {
	var comment Comment
	if err := db.DB.Model(&Comment{}).Where("comment_id = ?", commentId).First(&comment).Error; err != nil {
		log.Println("query comment by id failed, err:", err)
		return nil, err
	}
	log.Println("query comment by id sucess")
	return &comment, nil
}

// 获取全部根评论
func (CommentDao) GetRootCommentsByArticle(ArticleId int64) ([]*Comment, error) {
	var CommentList []*Comment
	if err := db.DB.Model(&Comment{}).Where("article_id = ? and parent_id is NULL", ArticleId).Find(&CommentList).Error; err != nil {
		log.Println("get root comments by article failed, err:", err)
		return nil, err
	}
	log.Println("get root comments by article success")
	return CommentList, nil
}

// GetRepliesByParentId 获取指定父评论的所有回复
func (CommentDao) GetRepliesByParentId(parentId int64) ([]*Comment, error) {
	var replies []*Comment
	if err := db.DB.Where("parent_id = ?", parentId).Find(&replies).Error; err != nil {
		log.Println("get replies by parent id failed, err:", err)
		return nil, err
	}
	log.Println("get replies by parent id success")
	return replies, nil
}

// RecursiveGetReplies 递归获取所有评论及其所有级别的回复
func (d CommentDao) RecursiveGetReplies(comment *Comment) error {
	replies, err := d.GetRepliesByParentId(comment.CommentId)
	if err != nil {
		return err
	}
	comment.Replies = replies

	// 递归获取每个回复的回复
	for i := range replies {
		if err := d.RecursiveGetReplies(replies[i]); err != nil {
			log.Println("get replies failed, err:", err)
			return err
		}
	}
	return nil
}
