package projects

type Project struct {
	ID     string `json:"projectId"`
	Name   string `json:"name"`
	Number string `json:"projectNumber"`
}

func (p *Project) String() string {
	// return fmt.Sprintf("%s (id: %s, number: %s)", p.Name, p.ID, p.Number)
	return p.Name
}

type Projects []Project
