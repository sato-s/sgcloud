package projects

type Project struct {
	ID     string `json:"projectId"`
	Name   string `json:"name"`
	Number string `json:"projectName"`
}

type Projects []Project
