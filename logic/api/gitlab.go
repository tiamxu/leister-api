package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tiamxu/kit/log"
	"github.com/tiamxu/leister-api/logic/service"
	"github.com/tiamxu/leister-api/types"
)

// GitlabHandler GitLab API 处理器
type GitlabHandler struct {
	service *service.GitlabService
}

// NewGitlabHandler 创建 GitLab 处理器实例
func NewGitlabHandler(service *service.GitlabService) *GitlabHandler {
	return &GitlabHandler{
		service: service,
	}
}

func (h *GitlabHandler) GetProject(c *gin.Context) {
	log.Infof("GetProject request received")
	var req types.GitlabProjectRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Errorf("GetProject request bind error: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	log.Infof("Getting GitLab project: %s in group: %s", req.Name, req.Group)
	resp, err := h.service.GetProject(c.Request.Context(), &req)
	if err != nil {
		log.Errorf("GetProject service error: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	log.Infof("GetProject request completed successfully")
	c.JSON(http.StatusOK, resp)
}

// GenProjects 生成 GitLab 项目数据
// @Summary 生成 GitLab 项目数据
// @Description 生成 GitLab 组下所有项目的数据
// @Tags gitlab
// @Accept json
// @Produce json
// @Param request body types.GitlabGenRequest true "GitLab 项目生成请求"
// @Success 200 {object} service.GitlabGenResponse
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/gitlab/gen [post]
func (h *GitlabHandler) GenProjects(c *gin.Context) {
	log.Infof("GenProjects request received")
	var req types.GitlabGenRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Errorf("GenProjects request bind error: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	log.Infof("Generating GitLab projects for group: %s", req.Group)
	resp, err := h.service.GenProjects(c.Request.Context(), &req)
	if err != nil {
		log.Errorf("GenProjects service error: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	log.Infof("GenProjects request completed successfully")
	c.JSON(http.StatusOK, resp)
}
