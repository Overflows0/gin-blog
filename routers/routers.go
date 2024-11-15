package routers

import (
	"gin-blog/pkg/setting"

	"gin-blog/routers/api"
	v1 "gin-blog/routers/api/v1"

	"gin-blog/middleware/jwt"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	gin.SetMode(setting.RunMode)

	r.GET("/auth", api.GetAuth)
	apiv1 := r.Group("/api/v1")
	apiv1.Use(jwt.JWT())
	{
		// 获取标签列表
		apiv1.GET("/tags", v1.GetTags)
		// 新建标签
		apiv1.POST("/tags", v1.AddTag)
		// 编辑指定标签
		apiv1.PUT("/tags/:id", v1.EditTag)
		// 删除指定标签
		apiv1.DELETE("/tags/:id", v1.DeleteTag)

		// 获取指定文章
		apiv1.GET("/articles/:id", v1.GetArticle)
		// 获取文章列表
		apiv1.GET("/articles", v1.GetArticles)
		// 新增文章
		apiv1.POST("/articles", v1.AddArticle)
		// 修改指定文章
		apiv1.PUT("/articles/:id", v1.EdditArticle)
		// 删除指定文章
		apiv1.DELETE("/articles/:id", v1.DeleteArticle)
	}

	return r
}
