package projects

type Project struct {
	ID     string `json:"projectId"`
	Name   string `json:"name"`
	Number string `json:"projectName"`
}

func (p *Project) String() string {
	return p.ID + p.Name + p.Number
}

type Projects []Project
