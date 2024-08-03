package main

import (
	"context"
	"fmt"
	"os"
	"os/exec"

	resourcemanager "cloud.google.com/go/resourcemanager/apiv3"
	resourcemanagerpb "cloud.google.com/go/resourcemanager/apiv3/resourcemanagerpb"
	"google.golang.org/api/iterator"

	"github.com/pterm/pterm"
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

	ctx := context.Background()
	c, err := resourcemanager.NewProjectsClient(ctx)
	if err != nil {
		pterm.Debug.Println(err)
		pterm.Error.Println("Failed to find default credentials. Try `gcloud auth application-default login`")
		os.Exit(1)
	}
	defer c.Close()

	// Search使ったほうがよさそう
	// https://github.com/googleapis/google-cloud-go/blob/3a566ed3f464089af85ab938bc593f2acb14fdf7/internal/generated/snippets/resourcemanager/apiv3/ProjectsClient/SearchProjects/main.go
	req := &resourcemanagerpb.ListProjectsRequest{
		// TODO: Fill request struct fields.
	}
	it := c.ListProjects(ctx, req)
	for {
		resp, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			handleFatalError(err)
		}
		pterm.Debug.Println(resp)
		_ = it.Response.(*resourcemanagerpb.ListProjectsResponse)
	}

	// use(path)

	// Initialize an empty slice to hold the options
	var options []string

	// Generate 100 options and add them to the options slice
	for i := 0; i < 100; i++ {
		options = append(options, fmt.Sprintf("Option %d", i))
	}

	// Generate 5 additional options with a specific message and add them to the options slice
	for i := 0; i < 5; i++ {
		options = append(options, fmt.Sprintf("You can use fuzzy searching (%d)", i))
	}

	// Use PTerm's interactive select feature to present the options to the user and capture their selection
	// The Show() method displays the options and waits for the user's input
	selectedOption, _ := pterm.DefaultInteractiveSelect.WithOptions(options).Show()

	// Display the selected option to the user with a green color for emphasis
	pterm.Info.Printfln("Selected option: %s", pterm.Green(selectedOption))
}
