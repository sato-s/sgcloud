package command

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"

	"github.com/pterm/pterm"
	"github.com/sato-s/sgcloud/internal/projects"
)

const defaultFilter = "NOT (projectId:sys-* OR projectId:qwiklabs-*)"

func ProjectList() (projects.Projects, error) {
	filter, ok := os.LookupEnv("SGCLOUD_PROJECT_FILTER")
	var args []string
	if ok {
		args = []string{"--format", "json", "projects", "list", "--filter", filter}
	} else {
		args = []string{"--format", "json", "projects", "list", "--filter", defaultFilter}
	}

	out, err := runGcloud(args...)

	if err != nil {
		return projects.Projects{}, err
	} else {
		pjs := make(projects.Projects, 0)
		err := json.Unmarshal(out, &pjs)
		if err != nil {
			return projects.Projects{}, fmt.Errorf("Failed to parse output. %v", err)
		} else {
			if len(pjs) == 0 {
				return nil, fmt.Errorf("No projects found: args=%v", args)
			}
			return pjs, nil
		}
	}
}

func ActivateSgcloudConfig(configName string) error {
	_, err := runGcloud("config", "configurations", "activate", configName)
	return err
}

func CreateSgcloudConfig(configName string) error {
	// This command create and activate `configName`
	_, err := runGcloud("config", "configurations", "create", configName)
	return err
}

func SetProject(id string) error {
	_, err := runGcloud("config", "set", "project", id, "--quiet")
	return err
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
