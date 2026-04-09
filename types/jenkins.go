package types

// JenkinsJobRequest Jenkins 任务创建请求
type JenkinsJobRequest struct {
	Name  string `json:"name"`
	Group string `json:"group"`
}

// JenkinsJobResponse Jenkins 任务创建响应
type JenkinsJobResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

type AddJobReq struct {
	Name  string `json:"name"`
	Group string `json:"group"`
}
