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

	selectedPjStr, err := pterm.
		DefaultInteractiveSelect.
		WithDefaultText("Select a Project").
		WithOptions(options).
		Show()
	if err != nil {
		pterm.Error.Println(err)
		os.Exit(1)
	}

	// Find selected Project
	for i, option := range options {
		if option == selectedPjStr {
			selectedPj := pjs[i]
			pterm.Debug.Printfln("Selected pj: %+v", pterm.Green(selectedPj))
			if err := command.SetProject(selectedPj.ID); err != nil {
				pterm.Error.Printfln("Unable to change project. %s", err)
			} else {
				pterm.Info.Printfln("Switched to %s", selectedPjStr)
				return
			}
		}
	}
	// Never reach here
	pterm.Error.Printfln("Unable to find selected project %s", selectedPjStr)
	os.Exit(1)
}
