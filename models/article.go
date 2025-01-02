package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Article struct {
	Model

	TagID      int    `json:"tag_id" gorm:"index"`
	Tag        Tag    `json:"tag"`
	Title      string `json:"title"`
	Desc       string `json:"desc"`
	Content    string `json:"content"`
	CreatedBy  string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
	State      int    `json:"state"`
}

func (Article) TableName() string {
	return "article"
}

type ArticleRepository struct {
	Repository[Article]
}

func NewArticleRepository(db *gorm.DB) *ArticleRepository {
	res := &ArticleRepository{}
	res.SetDB(db)
	return res
}

// 创建文章之前自动调用
func (article *Article) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("CreatedOn", time.Now().Unix())

	return nil
}

// 修改文章之前自动调用
func (article *Article) BeforeUpdate(scope *gorm.Scope) error {
	scope.SetColumn("ModifiedOn", time.Now().Unix())

	return nil
}

func (a *ArticleRepository) ExistArticleByID(id int) bool {
	var article Article
	a.DB.Select("id").Where("id = ?", id).First(&article)

	return article.ID > 0
}

func (a *ArticleRepository) GetArticleTotal(maps interface{}) (count int) {
	a.DB.Model(&Article{}).Where(maps).Count(&count)
	return
}

func (a *ArticleRepository) GetArticles(pageNum int, pageSize int, maps interface{}) (articles []Article) {
	a.DB.Preload("Tag").Where(maps).Offset(pageNum).Limit(pageSize).Find(&articles)
	return
}

func (a *ArticleRepository) GetArticle(id int) (article Article) {
	a.DB.Preload("Tag").Where("id = ?", id).First(&article)
	return
}

func (a *ArticleRepository) AddArticle(data map[string]interface{}) bool {
	a.DB.Create(&Article{
		TagID:     data["tag_id"].(int),
		Title:     data["title"].(string),
		Desc:      data["desc"].(string),
		Content:   data["content"].(string),
		CreatedBy: data["created_by"].(string),
		State:     data["state"].(int),
	})
	return true
}

func (a *ArticleRepository) EditArticle(id int, data interface{}) bool {
	a.DB.Model(&Article{}).Where("id = ?", id).Update(data)
	return true
}

func (a *ArticleRepository) DeleteArticle(id int) bool {
	a.DB.Where("id = ?", id).Delete(Article{})
	return true
}
