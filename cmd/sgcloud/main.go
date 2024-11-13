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

var logger *pterm.Logger

func main() {
	debugPrintPtr := flag.Bool("debug", false, "Enable debug print")
	openBrowserPtr := flag.Bool("b", false, "Open Browser")
	flag.Parse()

	if *debugPrintPtr {
		logger = pterm.DefaultLogger.WithLevel(pterm.LogLevelDebug).WithTime(false)
	} else {
		logger = pterm.DefaultLogger.WithLevel(pterm.LogLevelInfo).WithTime(false)
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
		logger.Error("`gcloud` command not found.")
		os.Exit(1)
	}
	logger.Debug(path)
}

func getProjects() projects.Projects {
	cache, err1 := cache.NewCache()
	if err1 != nil {
		logger.Debug("Failed to read cache.", logger.Args("err", err1))
	}
	if cache.Projects != nil && err1 == nil && !(cache.IsExpired()) {
		return cache.Projects
	} else {
		spinnerInfo, _ := pterm.DefaultSpinner.Start("Getting project list...")
		pjs, err2 := command.ProjectList()
		if err2 != nil {
			logger.Error("Failed to get project list.", logger.Args("err", err2))
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
		logger.Error("Failed to run selector", logger.Args("err", err))
		os.Exit(1)
	}

	// Find selected Project
	for i, option := range options {
		if option == selectedPjStr {
			return pjs[i]
		}
	}
	// Never reach here
	logger.Error("Unable to find selected project", logger.Args("selectedPjStr", selectedPjStr))
	os.Exit(1)
	return projects.Project{}
}

func gcloudProjectChange(selectedPj projects.Project) {
	logger.Debug("project selected", logger.Args("selectedPj", selectedPj))
	spinner, _ := pterm.DefaultSpinner.Start("gcloud config set project " + pterm.Blue(selectedPj.ID))
	if err := command.SetProject(selectedPj.ID); err != nil {
		spinner.Fail("Unable to change project.", logger.Args("err", err))
		os.Exit(1)
	} else {
		spinner.Info("Switched successfully")
		bulletListItems := []pterm.BulletListItem{
			{Level: 0, Text: "ID: " + pterm.Blue(selectedPj.ID)},
			{Level: 0, Text: "project name: " + pterm.Blue(selectedPj.Name)},
			{Level: 0, Text: "project number: " + pterm.Blue(selectedPj.Number)},
		}
		pterm.DefaultBulletList.WithItems(bulletListItems).Render()
		return
	}

}

func openBrowser(selectedPj projects.Project) {
	url := fmt.Sprintf("https://console.cloud.google.com/welcome?project=%s", selectedPj.ID)
	logger.Info("Opening project in browser", logger.Args("project", selectedPj.String()))
	err := browser.OpenBrowser(url)

	if err != nil {
		logger.Error("Failed to open browser", logger.Args("err", err))
		os.Exit(1)
	}
}
