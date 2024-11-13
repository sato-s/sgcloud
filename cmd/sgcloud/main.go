package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"

	"github.com/pterm/pterm"
	"github.com/sato-s/sgcloud/internal/browser"
	"github.com/sato-s/sgcloud/internal/cache"
	"github.com/sato-s/sgcloud/internal/command"
	"github.com/sato-s/sgcloud/internal/projects"
)

const sgCloudDefaultConfigName = "sgcloud"

func main() {
	debugPrintPtr := flag.Bool("debug", false, "Enable debug print")
	openBrowserPtr := flag.Bool("b", false, "Open Browser")
	flag.Parse()

	if *debugPrintPtr {
		pterm.EnableDebugMessages()
	}

	ensureGcloudInstalled()
	pjs := getProjects()
	selectedPj := showProjectSelector(pjs)
	if *openBrowserPtr {
		openBrowser(selectedPj)
	} else {
		gcloudProjectChange(selectedPj)
	}
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
	if cache.Projects != nil && err1 == nil && !(cache.IsExpired()) {
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

func showProjectSelector(pjs projects.Projects) projects.Project {
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
			return pjs[i]
		}
	}
	// Never reach here
	pterm.Error.Printfln("Unable to find selected project %s", selectedPjStr)
	os.Exit(1)
	return projects.Project{}
}

func gcloudProjectChange(selectedPj projects.Project) {
	pterm.Debug.Printfln("Selected pj: %+v", pterm.Green(selectedPj))
	if err := command.SetProject(selectedPj.ID); err != nil {
		pterm.Error.Printfln("Unable to change project. %s", err)
		os.Exit(1)
	} else {
		pterm.Success.Printfln("Switched to %s", selectedPj.Description())
		return
	}
}

func openBrowser(selectedPj projects.Project) {
	url := fmt.Sprintf("https://console.cloud.google.com/welcome?project=%s", selectedPj.ID)
	pterm.Success.Printfln("Opening %s in browser", selectedPj.String())
	err := browser.OpenBrowser(url)

	if err != nil {
		pterm.Error.Println("Failed to open browser", err)
		os.Exit(1)
	}
}
