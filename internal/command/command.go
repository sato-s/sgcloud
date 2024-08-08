package command

import (
	"encoding/json"
	"fmt"
	"os/exec"

	"github.com/pterm/pterm"
	"github.com/sato-s/sgcloud/internal/projects"
)

func ProjectList() (projects.Projects, error) {
	out, err := runGcloud("--format", "json", "projects", "list")
	if err != nil {
		return projects.Projects{}, err
	} else {
		pjs := make(projects.Projects, 0)
		err := json.Unmarshal(out, &pjs)
		if err != nil {
			pterm.Error.Println(string(out))
			return projects.Projects{}, fmt.Errorf("Failed to parse output. %v", err)
		} else {
			return pjs, nil
		}
	}
}

func runGcloud(args ...string) ([]byte, error) {
	out, err := exec.Command("gcloud", args...).CombinedOutput()

	if err != nil {
		pterm.Error.Println(string(out))
		return []byte{}, err
	} else {
		return out, nil
	}
}
