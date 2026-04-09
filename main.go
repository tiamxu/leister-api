package main

import (
	"github.com/gin-gonic/gin"
	conf "github.com/tiamxu/leister-api/config"
	"github.com/tiamxu/leister-api/api"
	"github.com/tiamxu/leister-api/routes"
	"github.com/tiamxu/leister-api/service"

	"github.com/tiamxu/kit/log"
)

func main() {
	// 加载配置
	cfg := conf.Load()

	// 先初始化配置（包括日志）
	if err := cfg.Initial(); err != nil {
		log.Fatalf("Config initialization failed: %v", err)
	}

	// 初始化服务
	gitlabSvc := service.NewGitlabService(cfg.GitLab)
	jenkinsSvc := service.NewJenkinsService(cfg.Jenkins)

	handlers := &api.Handlers{
		Gitlab:  api.NewGitlabHandler(gitlabSvc),
		Jenkins: api.NewJenkinsHandler(jenkinsSvc),
	}

	// 再创建 Gin 引擎（此时日志已初始化）
	r := gin.Default()
	routes.InitRoutes(r, handlers)
	if err := r.Run(cfg.HttpSrv.Address); err != nil {
		log.Fatalf("服务启动失败: %v", err)
	}

}
