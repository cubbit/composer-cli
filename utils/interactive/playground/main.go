package main

import (
	"fmt"
	"time"

	"github.com/cubbit/composer-cli/utils/interactive"
	"github.com/cubbit/composer-cli/utils/interactive/tui/input"
)

func main() {
	ic := interactive.New(interactive.Config{
		Steps: []string{
			"Project info",
			"Region",
			"Deploying",
			"Configuring",
			"Summary",
		},
	})

	name, err := ic.Input("Enter project name",
		input.WithPlaceholder("my-project"),
		input.WithRequired(),
	)
	if err != nil {
		ic.Error("Cancelled")
		return
	}
	ic.Success("Name: " + name)
	ic.NextStep()

	region, err := ic.Select("Select deployment region", []string{
		"EU-West",
		"US-East",
		"APAC",
	})
	if err != nil {
		ic.Error("Cancelled")
		return
	}
	ic.Success("Region: " + region)

	features, err := ic.MultiSelect("Select features", []string{
		"Monitoring",
		"Auto-scaling",
		"Backups",
		"Logging",
	})
	if err != nil {
		ic.Error("Cancelled")
		return
	}
	ic.Success(fmt.Sprintf("Features: %v", features))
	ic.NextStep()

	spinH := ic.StartSpinner("Deploying infrastructure...")
	for i := 0; i < 20; i++ {
		spinH.SetText(fmt.Sprintf("Deploying infrastructure... step %d/20", i+1))
		time.Sleep(100 * time.Millisecond)
	}
	spinH.Stop()
	ic.Success("Infrastructure deployed")
	ic.NextStep()

	progH := ic.StartProgress("Configuring nodes...")
	for i := 1; i <= 10; i++ {
		progH.Set(float64(i)/10.0, fmt.Sprintf("Node %d/10", i))
		time.Sleep(200 * time.Millisecond)
	}
	progH.Stop()
	ic.Success("Configuration complete")
	ic.NextStep()

	statusResult, err := interactive.Observe(ic, "Syncing data...", func(update func(string)) (string, error) {
		update("Connecting to remote...")
		time.Sleep(500 * time.Millisecond)
		update("Downloading metadata...")
		time.Sleep(500 * time.Millisecond)
		update("Verifying checksums...")
		time.Sleep(500 * time.Millisecond)
		return "sync-complete", nil
	})
	if err != nil {
		ic.Error("Sync failed")
		return
	}
	ic.Success("Sync result: " + statusResult)

	confirmed, err := ic.Confirm("Deploy project " + name + " to " + region + "?")
	if err != nil {
		ic.Error("Cancelled")
		return
	}
	if !confirmed {
		ic.Info("Deployment skipped")
		return
	}

	err = ic.Spin("Running final deployment...", func() error {
		time.Sleep(1500 * time.Millisecond)
		return nil
	})
	if err != nil {
		ic.Error("Deployment failed")
		return
	}

	ic.Success("Project deployed successfully!")
	ic.Info("You can monitor it in the dashboard")
}
