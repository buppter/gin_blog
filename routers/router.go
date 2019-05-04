package routers

import (
	"github.com/gin-gonic/gin"
	"go_code/gin-blog/routers/api"
	v1 "go_code/gin-blog/routers/api/v1"

	"go_code/gin-blog/middleware/jwt"
	"go_code/gin-blog/pkg/setting"
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
		//获取标签列表
		apiv1.GET("/tags", v1.GetTags)

		//新建标签
		apiv1.POST("/tags", v1.AddTag)

		//更新指定标签
		apiv1.PUT("/tags/:id", v1.EditTag)

		//删除指定标签
		apiv1.DELETE("/tags/:id", v1.DeleteTag)

		//get the article list
		apiv1.GET("/articles", v1.GetArticles)

		//get target article
		apiv1.GET("/articles/:id", v1.GetArticle)

		//add new article
		apiv1.POST("/articles", v1.AddArticle)

		//update article
		apiv1.PUT("/articles/:id", v1.EditArticle)

		//delete article
		apiv1.DELETE("/articles/:id", v1.DeleteArticle)
	}

	return r
}
