package command

import (
	"os/exec"

	"github.com/pterm/pterm"
)

func ProjectList() error {
	out, err := runGcloud("--format", "json", "projects", "list")
	if err != nil {
		return err
	} else {
		pterm.Info.Println(string(out))
	}

	return nil
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
