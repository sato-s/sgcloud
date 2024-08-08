package main

import (
	"os"
	"os/exec"

	"github.com/pterm/pterm"
	"github.com/sato-s/sgcloud/internal/command"
)

func main() {

	pterm.EnableDebugMessages()

	path, err := exec.LookPath("gcloud")
	if err != nil {
		pterm.Error.Println("`gcloud` command not found.")
		os.Exit(1)
	}
	pterm.Debug.Println(path)
	pjs, err := command.ProjectList()
	if err != nil {
		pterm.Error.Println(err)
		os.Exit(1)
	}

	var options []string
	for _, pj := range pjs {
		options = append(options, pj.String())
	}

	selectedOption, _ := pterm.
		DefaultInteractiveSelect.
		WithDefaultText("Select a Project").
		WithOptions(options).
		Show()

	pterm.Info.Printfln("Selected option: %s", pterm.Green(selectedOption))
}
