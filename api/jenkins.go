package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tiamxu/kit/log"
	"github.com/tiamxu/leister-api/service"
	"github.com/tiamxu/leister-api/types"
)

// JenkinsHandler Jenkins API 处理器
type JenkinsHandler struct {
	service *service.JenkinsService
}

// NewJenkinsHandler 创建 Jenkins 处理器实例
func NewJenkinsHandler(service *service.JenkinsService) *JenkinsHandler {
	return &JenkinsHandler{
		service: service,
	}
}


func (h *JenkinsHandler) CreateJob(c *gin.Context) {
	log.Infof("CreateJob request received")
	var req types.JenkinsJobRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Errorf("CreateJob request bind error: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	log.Infof("Creating Jenkins job: %s", req.Name)
	resp, err := h.service.CreateJob(c.Request.Context(), &req)
	if err != nil {
		log.Errorf("CreateJob service error: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	log.Infof("CreateJob request completed successfully")
	c.JSON(http.StatusOK, resp)
}


func (h *JenkinsHandler) CreateJobs(c *gin.Context) {
	log.Infof("CreateJobs request received")
	var reqs []*types.JenkinsJobRequest
	if err := c.ShouldBindJSON(&reqs); err != nil {
		log.Errorf("CreateJobs request bind error: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	log.Infof("Creating %d Jenkins jobs", len(reqs))
	resp, err := h.service.CreateJobs(c.Request.Context(), reqs)
	if err != nil {
		log.Errorf("CreateJobs service error: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	log.Infof("CreateJobs request completed successfully")
	c.JSON(http.StatusOK, resp)
}

// UpdateJob 更新 Jenkins 任务
// @Summary 更新 Jenkins 任务
// @Description 更新 Jenkins 任务配置
// @Tags jenkins
// @Accept json
// @Produce json
// @Param request body types.JenkinsJobRequest true "Jenkins 任务更新请求"
// @Success 200 {object} service.JenkinsJobResponse
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/jenkins/update [post]
func (h *JenkinsHandler) UpdateJob(c *gin.Context) {
	log.Infof("UpdateJob request received")
	var req types.JenkinsJobRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Errorf("UpdateJob request bind error: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	log.Infof("Updating Jenkins job: %s", req.Name)
	resp, err := h.service.UpdateJob(c.Request.Context(), &req)
	if err != nil {
		log.Errorf("UpdateJob service error: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	log.Infof("UpdateJob request completed successfully")
	c.JSON(http.StatusOK, resp)
}
