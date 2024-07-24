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

// func handleFatalError(error Error) {
// }
func main() {

	pterm.EnableDebugMessages()

	path, err := exec.LookPath("gcloud")
	if err != nil {
		pterm.Error.Println("gcloud not found.")
		os.Exit(1)
	}
	pterm.Debug.Println(path)

	ctx := context.Background()
	c, err := resourcemanager.NewProjectsClient(ctx)
	if err != nil {
		// TODO: Handle error.
		panic(err)
	}
	defer c.Close()

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
			panic(err)
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
