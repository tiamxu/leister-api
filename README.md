# Leister API

Leister API 是 Leister DevOps 工具集的后端服务，提供 Jenkins 任务管理和 GitLab 项目管理的 RESTful API 接口。

## 功能特性

### 核心功能

1. **Jenkins 任务管理**
   - 创建单个 Jenkins 任务
   - 批量创建 Jenkins 任务
   - 更新 Jenkins 任务配置

2. **GitLab 项目管理**
   - 获取 GitLab 项目信息
   - 生成 GitLab 组项目数据

## 项目结构

```
leister-api/
├── api/                 # API 处理层
│   ├── gitlab.go        # GitLab 处理器
│   ├── handler.go       # 处理器通用结构
│   └── jenkins.go       # Jenkins 处理器
├── config/              # 配置管理
│   ├── config.go        # 配置结构定义
│   └── config.yaml      # 配置文件
├── model/               # 数据模型
│   └── item.go          # 项目数据模型
├── repo/                # 数据访问层
│   └── init.go          # 数据库初始化
├── routes/              # 路由层
│   └── router.go        # 路由设置
├── service/             # 业务逻辑层
│   ├── gitlab.go        # GitLab 服务
│   └── jenkins.go       # Jenkins 服务
├── types/               # 类型定义
│   ├── gitlab.go        # GitLab 相关类型
│   └── jenkins.go       # Jenkins 相关类型
├── main.go              # 主入口
├── go.mod               # 依赖管理
└── README.md            # 文档
```

## 分层架构

Leister API 采用清晰的分层架构，参考 cactus 项目的设计思路：

### 1. 表现层（API）
- **职责**：处理 HTTP 请求、参数校验、响应格式化
- **文件**：`api/` 目录下的处理器
- **依赖**：依赖服务层（Service）

### 2. 业务层（Service）
- **职责**：实现业务逻辑、处理外部服务调用
- **文件**：`service/` 目录下的服务
- **依赖**：依赖数据访问层（Repo）

### 3. 数据访问层（Repo）
- **职责**：与数据库交互、数据持久化
- **文件**：`repo/` 目录下的数据访问
- **依赖**：无（只依赖数据库驱动）

### 4. 路由层（Routes）
- **职责**：注册路由、中间件配置
- **文件**：`routes/router.go`
- **依赖**：依赖表现层（API）

### 5. 配置层（Config）
- **职责**：加载和管理配置
- **文件**：`config/` 目录
- **依赖**：无

## 安装

### 构建项目

```bash
go build -o leister-api
```

### 安装到系统路径

```bash
# Linux/Mac
cp leister-api /usr/local/bin/
chmod +x /usr/local/bin/leister-api

# Windows
copy leister-api.exe C:\Windows\System32\
```

## 配置

Leister API 使用 `config.yaml` 配置文件。

### 配置文件示例

```yaml
httpSrv:
  address: :8080
  timeout: 30

jenkins:
  url: https://jenkins.example.com
  username: admin
  password: admin-password

gitlab:
  url: https://gitlab.example.com
  token: your-token

db:
  driver: mysql
  host: localhost
  port: 3306
  database: leister
  username: root
  password: password
  maxIdleConns: 10
  maxOpenConns: 100
  connMaxLifetime: 60

log:
  level: info
  type: stdout
  format: text
```

### 配置项说明

| 配置项 | 说明 | 默认值 |
|-------|------|--------|
| `httpSrv.address` | 服务监听地址 | `:8080` |
| `httpSrv.timeout` | 请求超时时间（秒） | `30` |
| `jenkins.url` | Jenkins 服务地址 | - |
| `jenkins.username` | Jenkins 用户名 | - |
| `jenkins.password` | Jenkins 密码 | - |
| `gitlab.url` | GitLab 服务地址 | - |
| `gitlab.token` | GitLab 访问令牌 | - |
| `db.*` | 数据库配置 | - |
| `log.*` | 日志配置 | - |

## 使用指南

### 启动服务

```bash
# 使用默认配置文件（config.yaml）
./leister-api

# 指定配置文件
./leister-api -c /path/to/config.yaml
```

### 健康检查

```bash
curl http://localhost:8080/health
```

## API 接口

### Jenkins 接口

#### 创建 Jenkins 任务

```bash
POST /api/jenkins/create
Content-Type: application/json

{
  "name": "job-name",
  "group": "job-group"
}
```

响应：
```json
{
  "status": "success",
  "message": "Job job-name created successfully"
}
```

#### 批量创建 Jenkins 任务

```bash
POST /api/jenkins/cts
Content-Type: application/json

{
  "group": "job-group"
}
```

响应：
```json
{
  "status": "success",
  "message": "Created 3 jobs successfully"
}
```

#### 更新 Jenkins 任务

```bash
POST /api/jenkins/update
Content-Type: application/json

{
  "name": "job-name",
  "group": "job-group"
}
```

响应：
```json
{
  "status": "success",
  "message": "Job job-name updated successfully"
}
```

### GitLab 接口

#### 获取 GitLab 项目信息

```bash
POST /api/gitlab/project
Content-Type: application/json

{
  "name": "project-name",
  "group": "project-group"
}
```

响应：
```json
{
  "status": "success",
  "project": {
    "id": 123,
    "name": "project-name",
    "group": "project-group",
    "http_url_to_repo": "https://gitlab.example.com/group/project-name.git",
    "ssh_url_to_repo": "git@gitlab.example.com:group/project-name.git"
  }
}
```

#### 生成 GitLab 项目数据

```bash
POST /api/gitlab/gen
Content-Type: application/json

{
  "group": "project-group"
}
```

响应：
```json
{
  "status": "success",
  "message": "Generated 5 projects",
  "projects": [
    {
      "id": 1,
      "name": "project1",
      "group": "project-group",
      "http_url_to_repo": "https://gitlab.example.com/group/project1.git",
      "ssh_url_to_repo": "git@gitlab.example.com:group/project1.git"
    }
  ]
}
```

## 技术栈

- **开发语言**：Go 1.25+
- **Web 框架**：Gin（通过 kit/http 包装）
- **日志库**：kit/log
- **依赖库**：
  - github.com/gin-gonic/gin - Web 框架
  - github.com/bndr/gojenkins - Jenkins 客户端
  - github.com/xanzy/go-gitlab - GitLab 客户端
  - github.com/koding/multiconfig - 配置管理
  - github.com/tiamxu/kit - 工具库

## 开发

### 环境要求

- Go 1.25 或更高版本
- Jenkins 服务（用于 Jenkins 功能）
- GitLab 服务（用于 GitLab 功能）
- MySQL 数据库（用于存储项目数据）

### 项目开发

1. 克隆项目
2. 配置 config.yaml
3. 安装依赖：`go mod tidy`
4. 构建项目：`go build -o leister-api`
5. 运行测试：`go test ./...`

## 与 leister CLI 配合使用

Leister API 需要与 leister CLI 配合使用：

```bash
# 1. 启动 leister-api 服务
./leister-api

# 2. 在另一个终端使用 leister CLI
cd ../leister
./gigctl jks create -n myjob -g mygroup
./gigctl git get -n myproject -g mygroup
```

## 贡献指南

欢迎贡献代码、报告问题和提出建议！

## 许可证

本项目采用 MIT 许可证。
