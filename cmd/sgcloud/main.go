package main

import (
	"os"
	"os/exec"

	"github.com/pterm/pterm"
	"github.com/sato-s/sgcloud/internal/command"
)

func handleFatalError(err error) {
	pterm.Error.Println(err)
	os.Exit(1)
}

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
	// use(path)

	// Initialize an empty slice to hold the options
	var options []string
	for _, pj := range pjs {
		options = append(options, pj.String())
	}

	// Use PTerm's interactive select feature to present the options to the user and capture their selection
	// The Show() method displays the options and waits for the user's input
	selectedOption, _ := pterm.DefaultInteractiveSelect.WithOptions(options).Show()

	// Display the selected option to the user with a green color for emphasis
	pterm.Info.Printfln("Selected option: %s", pterm.Green(selectedOption))
}
