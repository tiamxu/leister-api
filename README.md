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
├── config/              # 配置管理
│   ├── config.go        # 配置结构定义
│   └── config.yaml      # 配置文件
├── api/                 # API 相关
│   ├── handler/         # 请求处理器
│   │   ├── jenkins.go   # Jenkins 处理器
│   │   └── gitlab.go    # GitLab 处理器
│   └── router.go        # 路由配置
├── service/             # 业务服务
│   ├── jenkins/         # Jenkins 服务
│   │   └── service.go   # Jenkins 服务实现
│   └── gitlab/          # GitLab 服务
│       └── service.go   # GitLab 服务实现
├── main.go              # 主入口
├── go.mod               # 依赖管理
└── README.md            # 文档
```

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
server:
  addr: :8080
  timeout: 30

jenkins:
  url: https://jenkins.example.com
  username: admin
  password: admin-password

gitlab:
  url: https://gitlab.example.com
  token: your-token

log:
  level: info
  type: stdout
```

### 配置项说明

| 配置项 | 说明 | 默认值 |
|-------|------|--------|
| `server.addr` | 服务监听地址 | `:8080` |
| `server.timeout` | 请求超时时间（秒） | `30` |
| `jenkins.url` | Jenkins 服务地址 | - |
| `jenkins.username` | Jenkins 用户名 | - |
| `jenkins.password` | Jenkins 密码 | - |
| `gitlab.url` | GitLab 服务地址 | - |
| `gitlab.token` | GitLab 访问令牌 | - |
| `log.level` | 日志级别 | `info` |
| `log.type` | 日志输出类型 | `stdout` |

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

[
  {"name": "job1", "group": "group1"},
  {"name": "job2", "group": "group1"},
  {"name": "job3", "group": "group1"}
]
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
- **Web 框架**：Gin
- **日志库**：kit/log
- **依赖库**：
  - github.com/gin-gonic/gin - Web 框架
  - github.com/bndr/gojenkins - Jenkins 客户端
  - github.com/xanzy/go-gitlab - GitLab 客户端
  - github.com/koding/multiconfig - 配置管理
  - github.com/tiamxu/kit - 工具库

## 架构说明

### 分层架构

Leister API 采用分层架构设计：

- **API 层**（handler）：负责 HTTP 请求处理、参数校验、响应格式化
- **服务层**（service）：负责业务逻辑、外部服务集成
- **配置层**（config）：负责配置加载和管理

### 请求流程

```
HTTP Request -> Router -> Handler -> Service -> External API (Jenkins/GitLab)
                                    |
                                    v
HTTP Response <- Handler <- Service
```

## 开发

### 环境要求

- Go 1.25 或更高版本
- Jenkins 服务（用于 Jenkins 功能）
- GitLab 服务（用于 GitLab 功能）

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

## 部署

### Docker 部署

```dockerfile
FROM golang:1.25-alpine AS builder
WORKDIR /app
COPY . .
RUN go build -o leister-api

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /app/leister-api .
COPY config.yaml .
CMD ["./leister-api"]
```

构建并运行：

```bash
docker build -t leister-api .
docker run -p 8080:8080 -v /path/to/config.yaml:/root/config.yaml leister-api
```

### Systemd 服务

创建 `/etc/systemd/system/leister-api.service`：

```ini
[Unit]
Description=Leister API Service
After=network.target

[Service]
Type=simple
User=root
WorkingDirectory=/opt/leister-api
ExecStart=/opt/leister-api/leister-api
Restart=on-failure

[Install]
WantedBy=multi-user.target
```

启用并启动服务：

```bash
systemctl enable leister-api
systemctl start leister-api
```

## 贡献指南

欢迎贡献代码、报告问题和提出建议！

## 许可证

本项目采用 MIT 许可证。
