package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	conf "github.com/tiamxu/leister-api/config"
	"github.com/tiamxu/leister-api/logic/api"
	"github.com/tiamxu/leister-api/logic/routes"
	"github.com/tiamxu/leister-api/logic/service"

	"github.com/tiamxu/kit/http"
	"github.com/tiamxu/kit/log"
)

func main() {
	// 加载配置
	cfg := conf.Load()
	if err := cfg.Initial(); err != nil {
		log.Fatalf("Config initialization failed: %v", err)
	}

	// 初始化服务
	jenkinsService := service.NewJenkinsService(cfg)
	gitlabService := service.NewGitlabService(cfg)

	// 初始化 API 处理器
	jenkinsHandler := api.NewJenkinsHandler(jenkinsService)
	gitlabHandler := api.NewGitlabHandler(gitlabService)

	// 设置路由
	router := http.NewGin(cfg.HttpSrv)
	routes.SetupRouter(router, jenkinsHandler, gitlabHandler)

	// 启动服务
	log.Infof("Server starting on %s", cfg.HttpSrv.Address)
	srv, err := http.StartServer(router, cfg.HttpSrv)
	if err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}

	// 等待中断信号
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	// 优雅关闭
	log.Infoln("Shutting down server...")
	_, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := http.ShutdownServer(srv); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}

	log.Infoln("Server exited")
}
