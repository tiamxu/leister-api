package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/tiamxu/leister-api/api"
)

// SetupRouter 设置路由
func InitRoutes(r *gin.Engine, h *api.Handlers) {
	// 健康检查
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "ok",
		})
	})

	// API 路由组
	api := r.Group("/api")
	{
		// Jenkins 路由
		jenkins := api.Group("/jenkins")
		{
			jenkins.POST("/create", h.Jenkins.CreateJob)
			jenkins.POST("/cts", h.Jenkins.CreateJobs)
			jenkins.POST("/update", h.Jenkins.UpdateJob)
		}

		// GitLab 路由
		gitlab := api.Group("/gitlab")
		{
			gitlab.POST("/project", h.Gitlab.GetProject)
			gitlab.POST("/gen", h.Gitlab.GenProjects)
		}
	}
}
