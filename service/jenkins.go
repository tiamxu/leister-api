package service

import (
	"context"
	"fmt"

	"github.com/bndr/gojenkins"
	"github.com/tiamxu/kit/log"
	"github.com/tiamxu/leister-api/config"
	"github.com/tiamxu/leister-api/types"
)

// JenkinsService Jenkins 服务
type JenkinsService struct {
	client *gojenkins.Jenkins
}

// NewJenkinsService 创建 Jenkins 服务实例
func NewJenkinsService(cfg config.JenkinsConfig) *JenkinsService {
	return &JenkinsService{
		client: gojenkins.CreateJenkins(nil, cfg.URL, cfg.Username, cfg.Password),
	}
}

// CreateJob 创建 Jenkins 任务
func (s *JenkinsService) CreateJob(ctx context.Context, req *types.JenkinsJobRequest) (*types.JenkinsJobResponse, error) {
	// 初始化 Jenkins 客户端
	log.Infof("Creating Jenkins job: %s", req.Name)
	_, err := s.client.Init(ctx)
	if err != nil {
		log.Errorf("Jenkins init error: %v", err)
		return nil, fmt.Errorf("jenkins init error: %v", err)
	}

	// 生成 Jenkins 任务配置
	configXML := s.generateJobConfig(req.Name, req.Group)

	// 检查任务是否存在
	exists, err := s.client.GetJob(ctx, req.Name)
	if err == nil && exists != nil {
		// 更新任务
		job := s.client.UpdateJob(ctx, configXML, req.Name)
		if job == nil {
			return nil, fmt.Errorf("jenkins update job error: job is nil")
		}
		return &types.JenkinsJobResponse{
			Status:  "success",
			Message: fmt.Sprintf("Job %s updated successfully", req.Name),
		}, nil
	}

	// 创建任务
	_, err = s.client.CreateJob(ctx, configXML, req.Name)
	if err != nil {
		return nil, fmt.Errorf("jenkins create job error: %v", err)
	}

	return &types.JenkinsJobResponse{
		Status:  "success",
		Message: fmt.Sprintf("Job %s created successfully", req.Name),
	}, nil
}

// CreateJobs 批量创建 Jenkins 任务
func (s *JenkinsService) CreateJobs(ctx context.Context, reqs []*types.JenkinsJobRequest) (*types.JenkinsJobResponse, error) {
	// 初始化 Jenkins 客户端
	_, err := s.client.Init(ctx)
	if err != nil {
		return nil, fmt.Errorf("jenkins init error: %v", err)
	}

	for _, req := range reqs {
		// 生成 Jenkins 任务配置
		configXML := s.generateJobConfig(req.Name, req.Group)

		// 检查任务是否存在
		exists, err := s.client.GetJob(ctx, req.Name)
		if err == nil && exists != nil {
			// 跳过已存在的任务
			continue
		}

		// 创建任务
		_, err = s.client.CreateJob(ctx, configXML, req.Name)
		if err != nil {
			return nil, fmt.Errorf("jenkins create job error for %s: %v", req.Name, err)
		}
	}

	return &types.JenkinsJobResponse{
		Status:  "success",
		Message: fmt.Sprintf("Created %d jobs successfully", len(reqs)),
	}, nil
}

// UpdateJob 更新 Jenkins 任务
func (s *JenkinsService) UpdateJob(ctx context.Context, req *types.JenkinsJobRequest) (*types.JenkinsJobResponse, error) {
	// 初始化 Jenkins 客户端
	_, err := s.client.Init(ctx)
	if err != nil {
		return nil, fmt.Errorf("jenkins init error: %v", err)
	}

	// 生成 Jenkins 任务配置
	configXML := s.generateJobConfig(req.Name, req.Group)

	// 更新任务
	job := s.client.UpdateJob(ctx, configXML, req.Name)
	if job == nil {
		return nil, fmt.Errorf("jenkins update job error: job is nil")
	}

	return &types.JenkinsJobResponse{
		Status:  "success",
		Message: fmt.Sprintf("Job %s updated successfully", req.Name),
	}, nil
}

// generateJobConfig 生成 Jenkins 任务配置
func (s *JenkinsService) generateJobConfig(name, group string) string {
	return fmt.Sprintf(`<?xml version='1.1' encoding='UTF-8'?>
<flow-definition plugin="workflow-job@1292.v27d8cc3e2602">
<actions>
  <org.jenkinsci.plugins.pipeline.modeldefinition.actions.DeclarativeJobAction plugin="pipeline-model-definition@2.2131.vb_9788088fdb_5"/>
  <org.jenkinsci.plugins.pipeline.modeldefinition.actions.DeclarativeJobPropertyTrackerAction plugin="pipeline-model-definition@2.2131.vb_9788088fdb_5">
	<jobProperties/>
	<triggers/>
	<parameters/>
	<options/>
  </org.jenkinsci.plugins.pipeline.modeldefinition.actions.DeclarativeJobPropertyTrackerAction>
</actions>
<description></description>
<keepDependencies>false</keepDependencies>
<properties/>
<definition class="org.jenkinsci.plugins.workflow.cps.CpsFlowDefinition" plugin="workflow-cps@3659.v582dc37621d8">
  <script>node{
	  stage(&apos;Loading&apos;)
	  def rootDir = pwd()
	  println(rootDir)
	  def pipeline = load &apos;pipeline.groovy&apos;
	  pipeline(&apos;%s&apos;,&apos;%s&apos;)
}</script>
  <sandbox>true</sandbox>
</definition>
<triggers/>
<disabled>false</disabled>
</flow-definition>`, name, group)
}
