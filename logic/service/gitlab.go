package service

import (
	"context"
	"fmt"

	"github.com/tiamxu/kit/log"
	"github.com/tiamxu/leister-api/config"
	"github.com/tiamxu/leister-api/types"
	"github.com/xanzy/go-gitlab"
)

// GitlabService GitLab 服务
type GitlabService struct {
	config *config.Config
	client *gitlab.Client
}

// NewGitlabService 创建 GitLab 服务实例
func NewGitlabService(cfg *config.Config) *GitlabService {
	client, err := gitlab.NewClient(cfg.GitLab.Token, gitlab.WithBaseURL(cfg.GitLab.URL))
	if err != nil {
		panic(fmt.Sprintf("Failed to create gitlab client: %v", err))
	}
	return &GitlabService{
		config: cfg,
		client: client,
	}
}

// GetProject 获取 GitLab 项目信息
func (s *GitlabService) GetProject(ctx context.Context, req *types.GitlabProjectRequest) (*types.GitlabProjectResponse, error) {
	log.Infof("Getting GitLab project: %s in group: %s", req.Name, req.Group)
	// 查找组
	groupOption := &gitlab.ListGroupsOptions{Search: gitlab.String(req.Group)}
	groups, _, err := s.client.Groups.ListGroups(groupOption, gitlab.WithContext(ctx))
	if err != nil {
		log.Errorf("Failed to get groups: %v", err)
		return nil, fmt.Errorf("failed to get groups: %v", err)
	}

	if len(groups) == 0 {
		log.Errorf("Group not found: %s", req.Group)
		return nil, fmt.Errorf("group not found: %s", req.Group)
	}

	group := groups[0]
	log.Infof("Found group: %s (ID: %d)", group.Name, group.ID)

	// 查找项目
	projectOption := &gitlab.ListGroupProjectsOptions{
		Search: gitlab.String(req.Name),
		ListOptions: gitlab.ListOptions{
			Page:    1,
			PerPage: 50,
		},
	}

	projects, _, err := s.client.Groups.ListGroupProjects(group.ID, projectOption, gitlab.WithContext(ctx))
	if err != nil {
		log.Errorf("Failed to get projects: %v", err)
		return nil, fmt.Errorf("failed to get projects: %v", err)
	}

	if len(projects) == 0 {
		log.Errorf("Project not found: %s", req.Name)
		return nil, fmt.Errorf("project not found: %s", req.Name)
	}

	project := projects[0]
	log.Infof("Found project: %s (ID: %d)", project.Name, project.ID)

	return &types.GitlabProjectResponse{
		Status: "success",
		Project: &types.ProjectInfo{
			ID:            project.ID,
			Name:          project.Name,
			Group:         req.Group,
			HTTPURLToRepo: project.HTTPURLToRepo,
			SSHURLToRepo:  project.SSHURLToRepo,
		},
	}, nil
}

// GenProjects 生成 GitLab 项目数据
func (s *GitlabService) GenProjects(ctx context.Context, req *types.GitlabGenRequest) (*types.GitlabGenResponse, error) {
	log.Infof("Generating GitLab projects for group: %s", req.Group)
	// 查找组
	groupOption := &gitlab.ListGroupsOptions{Search: gitlab.String(req.Group)}
	groups, _, err := s.client.Groups.ListGroups(groupOption, gitlab.WithContext(ctx))
	if err != nil {
		log.Errorf("Failed to get groups: %v", err)
		return nil, fmt.Errorf("failed to get groups: %v", err)
	}

	if len(groups) == 0 {
		log.Errorf("Group not found: %s", req.Group)
		return nil, fmt.Errorf("group not found: %s", req.Group)
	}

	group := groups[0]
	log.Infof("Found group: %s (ID: %d)", group.Name, group.ID)

	// 获取组下所有项目
	projectOption := &gitlab.ListGroupProjectsOptions{
		ListOptions: gitlab.ListOptions{
			Page:    1,
			PerPage: 50,
		},
	}

	projects, _, err := s.client.Groups.ListGroupProjects(group.ID, projectOption, gitlab.WithContext(ctx))
	if err != nil {
		log.Errorf("Failed to get projects: %v", err)
		return nil, fmt.Errorf("failed to get projects: %v", err)
	}

	log.Infof("Found %d projects in group: %s", len(projects), group.Name)

	// 构建响应
	projectInfos := make([]*types.ProjectInfo, 0, len(projects))
	for _, project := range projects {
		projectInfos = append(projectInfos, &types.ProjectInfo{
			ID:            project.ID,
			Name:          project.Name,
			Group:         req.Group,
			HTTPURLToRepo: project.HTTPURLToRepo,
			SSHURLToRepo:  project.SSHURLToRepo,
		})
	}

	log.Infof("Generated %d projects successfully", len(projectInfos))
	return &types.GitlabGenResponse{
		Status:   "success",
		Message:  fmt.Sprintf("Generated %d projects", len(projectInfos)),
		Projects: projectInfos,
	}, nil
}
