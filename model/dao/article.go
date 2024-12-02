package dao

import (
	"log"
	"our_blog/db"
	"sync"
	"time"
)

type Article struct {
	Title      string    `gorm:"column:title" json:"title"`             // 文章标题
	ArticleId  uint      `gorm:"column:article_id" json:"article_id"`   // 文章 ID
	AuthorId   uint      `gorm:"column:author_id" json:"author_id"`     // 作者 ID
	Excerpt    string    `gorm:"column:excerpt" json:"excerpt"`         // 文章摘要
	Category   string    `gorm:"column:category" json:"category"`       // 文章分类
	Content    string    `gorm:"column:content" json:"content"`         // 文章内容
	Status     string    `gorm:"column:status" json:"status"`           // 文章状态
	CreateTime time.Time `gorm:"column:create_time" json:"create_time"` // 创建时间
	UpdateTime time.Time `gorm:"column:update_time" json:"update_time"` // 更新时间
	// Tags     []Tag   `gorm:"many2many:article_tags;" json:"tags"` // 假设有一个标签表，并与文章是多对多关系
}

type ArticleDao struct {
}

var articleDao *ArticleDao
var articleOnce sync.Once

func NewArticleDaoInstance() *ArticleDao {
	articleOnce.Do(func() {
		articleDao = &ArticleDao{}
	})
	return articleDao
}

func (ArticleDao) CreateAnArticle(article *Article) (err error) {
	if err = db.DB.Create(article).Error; err != nil {
		log.Println("create an artilce failed, err : ", err)
		return err
	}
	return nil
}

func (ArticleDao) GetArticleById(articleId uint) (article Article, err error) {
	if err = db.DB.Where("article_id =?", articleId).First(&article).Error; err != nil {
		log.Println("get an article by id failed, err : ", err)
		return article, err
	}
	return article, nil
}

func (ArticleDao) GetAllArticle() (articles []Article, err error) {
	if err = db.DB.Find(&articles).Error; err != nil {
		log.Println("get all article failed, err : ", err)
		return articles, err
	}
	return articles, nil
}

func (ArticleDao) GetAricleListByPages(page, pageSize int) ([]Article, error) {
	var articles []Article
	// 计算跳过的记录数，基于当前页码和页面大小
	offset := (page - 1) * pageSize
	// 使用 Limit 和 Offset 方法实现分页
	if err := db.DB.Order("created_at desc").Limit(pageSize).Offset(offset).Find(&articles).Error; err != nil {
		log.Printf("find articles failed, err: %v", err)
		return nil, err
	}
	return articles, nil
}

func (ArticleDao) UpdateAnArticle(article *Article) (err error) {
	if err = db.DB.Model(article).Updates(article).Error; err != nil {
		log.Println("update an article failed, err: ", err)
		return err
	}
	return nil
}

func (ArticleDao) DeleteAnArticle(id int) (err error) {
	if err = db.DB.Where("article_id = ?", id).Delete(&Article{}).Error; err != nil {
		log.Println("delete an article failed, err: ", err)
		return err
	}
	return nil
}
