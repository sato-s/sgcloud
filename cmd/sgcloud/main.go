package main

import (
	"flag"
	"os"
	"os/exec"

	"github.com/pterm/pterm"
	"github.com/sato-s/sgcloud/internal/cache"
	"github.com/sato-s/sgcloud/internal/command"
	"github.com/sato-s/sgcloud/internal/projects"
)

const sgCloudDefaultConfigName = "sgcloud"

func main() {
	debugPrintPtr := flag.Bool("debug", false, "Enable debug print")
	flag.Parse()

	if *debugPrintPtr {
		pterm.EnableDebugMessages()
	}

	ensureGcloudInstalled()
	pjs := getProjects()
	showProjectSelector(pjs)
}

func ensureGcloudInstalled() {

	path, err := exec.LookPath("gcloud")
	if err != nil {
		pterm.Error.Println("`gcloud` command not found.")
		os.Exit(1)
	}
	pterm.Debug.Println(path)
}

func getProjects() projects.Projects {
	cache, err1 := cache.NewCache()
	if err1 != nil {
		pterm.Debug.Println("Failed to read cache.", err1)
	}
	if cache.Projects != nil && err1 == nil {
		return cache.Projects
	} else {
		spinnerInfo, _ := pterm.DefaultSpinner.Start("Getting project list...")
		pjs, err2 := command.ProjectList()
		if err2 != nil {
			pterm.Error.Println("Failed to get project list.", err2)
			os.Exit(1)
		}
		cache.Projects = pjs
		cache.Save()
		spinnerInfo.Info()
		return pjs
	}
}

func showProjectSelector(pjs projects.Projects) {
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
				os.Exit(1)
			} else {
				pterm.Success.Printfln("Switched to %s", selectedPjStr)
				return
			}
		}
	}
	// Never reach here
	pterm.Error.Printfln("Unable to find selected project %s", selectedPjStr)
	os.Exit(1)
}
