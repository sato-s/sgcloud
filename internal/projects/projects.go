package projects

import "fmt"

type Project struct {
	ID     string `json:"projectId"`
	Name   string `json:"name"`
	Number string `json:"projectNumber"`
}

func (p *Project) String() string {
	return p.Name
}

func (p *Project) Description() string {
	return fmt.Sprintf("%s (id: %s, number: %s)", p.Name, p.ID, p.Number)
}

type Projects []Project
