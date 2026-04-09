package config

import (
	"fmt"
	"os"

	"github.com/tiamxu/kit/log"
	"github.com/tiamxu/kit/sql"

	"github.com/koding/multiconfig"
	httpkit "github.com/tiamxu/kit/http"
	"github.com/tiamxu/leister-api/repo"
)

var (
	cfg        *Config
	name       = "leister-api"
	configPath = "config/config.yaml"
)

// Config 配置结构体
type Config struct {
	ENV     string               `yaml:"env"`
	Log     log.Config           `yaml:"log"`
	HttpSrv httpkit.ServerConfig `yaml:"http_srv"`
	DB      *sql.Config          `yaml:"db"`
	Jenkins JenkinsConfig        `yaml:"jenkins"`
	GitLab  GitLabConfig         `yaml:"gitlab"`
}

// JenkinsConfig Jenkins 配置
type JenkinsConfig struct {
	URL      string `yaml:"url"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

// GitLabConfig GitLab 配置
type GitLabConfig struct {
	URL   string `yaml:"url"`
	Token string `yaml:"token"`
}

// Initial 初始化配置
func (c *Config) Initial() (err error) {
	defer func() {
		if err == nil {
			log.Printf("config initialed, env: %s, name: %s", cfg.ENV, name)
		}
	}()

	// 初始化日志
	if err = log.InitLogger(&c.Log); err != nil {
		return fmt.Errorf("log init failed: %w", err)
	}

	// 数据库初始化
	if err = repo.Init(c.DB); err != nil {
		return fmt.Errorf("database init failed: %w", err)
	}

	return nil
}

// Load 加载配置
func Load() *Config {
	cfg = new(Config)

	env := os.Getenv("ENV")
	if env == "" {
		env = "dev"
	}

	switch env {
	case "dev":
		configPath = "config/config-dev.yaml"
	case "test":
		configPath = "config/config-test.yaml"
	case "prod":
		configPath = "config/config-prod.yaml"
	default:
		configPath = "config/config.yaml"
	}

	multiconfig.MustLoadWithPath(configPath, cfg)
	// 设置环境变量到配置
	cfg.ENV = env
	return cfg
}
