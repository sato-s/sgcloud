package projects

import "fmt"

type Project struct {
	ID     string `json:"projectId"`
	Name   string `json:"name"`
	Number string `json:"projectNumber"`
}

func (p *Project) String() string {
	return fmt.Sprintf("%s (name: %s, number: %s)", p.ID, p.Name, p.Number)
}

type Projects []Project
