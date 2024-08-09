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
			return projects.Projects{}, fmt.Errorf("Failed to parse output. %v", err)
		} else {
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
