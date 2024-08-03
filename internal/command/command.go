package command

import (
	"github.com/pterm/pterm"
	"os/exec"
)

func RunGcloud() error {
	out, err := exec.Command("gcloud", "projects", "list").Output()
	if err != nil {
		return err
	} else {
		pterm.Info.Println(string(out))
	}

	return nil
}
