package model

import (
	"log"
	"our_blog/dao"
	"sync"
	"time"
)

type Article struct {
	Title      string    `gorm:"column:title"`
	ArticleId  uint      `gorm:"column:article_id"`
	AuthorId   uint      `gorm:"column:author_id"`
	Excerpt    string    `gorm:"column:excerpt"`
	Category   string    `gorm:"column:category"`
	Content    string    `gorm:"column:content"`
	Status     string    `gorm:"column:status"`
	CreateTime time.Time `gorm:"column:craete_time"`
	UpdateTime time.Time `gorm:"column:update_time"`
	//Tags     []Tag   // 假设有一个标签表，并与文章是多对多关系
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
	if err = dao.DB.Create(article).Error; err != nil {
		log.Println("create an artilce failed, err : ", err)
		return err
	}
	return nil
}

func (ArticleDao) GetArticleById(articleId uint) (article Article, err error) {
	if err = dao.DB.Where("article_id =?", articleId).First(&article).Error; err != nil {
		log.Println("get an article by id failed, err : ", err)
		return article, err
	}
	return article, nil
}

func (ArticleDao) GetAllArticle() (articles []Article, err error) {
	if err = dao.DB.Find(&articles).Error; err != nil {
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
	if err := dao.DB.Order("created_at desc").Limit(pageSize).Offset(offset).Find(&articles).Error; err != nil {
		log.Printf("find articles failed, err: %v", err)
		return nil, err
	}
	return articles, nil
}

func (ArticleDao) UpdateAnArticle(article *Article) (err error) {
	if err = dao.DB.Model(article).Updates(article).Error; err != nil {
		log.Println("update an article failed, err: ", err)
		return err
	}
	return nil
}

func (ArticleDao) DeleteAnArticle(id int) (err error) {
	if err = dao.DB.Where("article_id = ?", id).Delete(&Article{}).Error; err != nil {
		log.Println("delete an article failed, err: ", err)
		return err
	}
	return nil
}
