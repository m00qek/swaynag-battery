package main

import (
	"encoding/json"
	"os/exec"
)

type Display struct {
	Name   string
	Active bool
}

func run() (string, error) {
	cmd := exec.Command("swaymsg", "-t", "get_outputs")

	output, err := cmd.Output()
	if err != nil {
		logError("Unable to get sway outputs.")
		return "", err
	}

	return string(output), nil
}

func filterActive(displays []Display) []string {
	var activeDisplays []string
	for _, display := range displays {
		if display.Active {
			activeDisplays = append(activeDisplays, display.Name)
		}
	}

	return activeDisplays
}

func ActiveDisplays() StringSet {
	var displays []Display
	jsonOutput, err := run()

	if err != nil {
		return EmptySet()
	}

	json.Unmarshal([]byte(jsonOutput), &displays)
	return SetFrom(filterActive(displays))
}
