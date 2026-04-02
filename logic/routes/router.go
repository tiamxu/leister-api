package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/tiamxu/leister-api/logic/api"
)

// SetupRouter 设置路由
func SetupRouter(r *gin.Engine, jenkinsHandler *api.JenkinsHandler, gitlabHandler *api.GitlabHandler) {
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
			jenkins.POST("/create", jenkinsHandler.CreateJob)
			jenkins.POST("/cts", jenkinsHandler.CreateJobs)
			jenkins.POST("/update", jenkinsHandler.UpdateJob)
		}

		// GitLab 路由
		gitlab := api.Group("/gitlab")
		{
			gitlab.POST("/project", gitlabHandler.GetProject)
			gitlab.POST("/gen", gitlabHandler.GenProjects)
		}
	}
}
