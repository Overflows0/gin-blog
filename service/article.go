package service

import (
	"gin-blog/models"
	"gin-blog/pkg/e"
	"gin-blog/pkg/setting"
)

type ArticleService struct {
	Repo models.ArticleRepository
}

// 获取单篇文章
func (a *ArticleService) GetArticle(id int) models.Article {
	a.Repo = *models.NewArticleRepository(models.ReturnDB())
	var data models.Article

	if a.Repo.ExistArticleByID(id) {
		data = a.Repo.GetArticle(id)
	} else {
		data.ID = -1
	}

	return data
}

// 获取多篇文章
func (a *ArticleService) GetArticles(maps map[string]interface{}) ([]models.Article, int) {
	a.Repo = *models.NewArticleRepository(models.ReturnDB())

	data := a.Repo.GetArticles(setting.PageNum, setting.PageSize, maps)
	total := a.Repo.GetArticleTotal(maps)

	return data, total
}

// 新增文章
func (a *ArticleService) AddArticle(maps map[string]interface{}) {

	a.Repo = *models.NewArticleRepository(models.ReturnDB())
	a.Repo.AddArticle(maps)

}

// 修改文章
func (a *ArticleService) EditArticle(maps map[string]interface{}) int {
	id := maps["id"].(int)
	tagId := maps["tag_id"].(int)

	code := e.INVALID_PARAMS
	a.Repo = *models.NewArticleRepository(models.ReturnDB())
	if a.Repo.ExistArticleByID(id) {
		if models.ExistTagByID(tagId) {

			a.Repo.EditArticle(id, maps)
			code = e.SUCCESS

		} else {
			code = e.ERROR_NOT_EXIST_TAG
		}
	} else {
		code = e.ERROR_NOT_EXIST_ARTICLE
	}

	return code
}

// 删除文章
func (a *ArticleService) DeleteArticle(id int) int {

	code := e.INVALID_PARAMS
	if a.Repo.ExistArticleByID(id) {
		a.Repo.DeleteArticle(id)
		code = e.SUCCESS
	} else {
		code = e.ERROR_NOT_EXIST_ARTICLE
	}

	return code
}
