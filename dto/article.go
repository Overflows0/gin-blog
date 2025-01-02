package dto

import "gin-blog/models"

type GetArticlesResponse struct {
	List  []models.Article `json:"list"`
	Total int              `json:"total"`
}
