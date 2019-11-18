package main

import (
	"encoding/json"
	"os/exec"
)

type Display struct {
	Name   string
	Active bool
}

func run() string {
	cmd := exec.Command("swaymsg", "-t", "get_outputs")
	output, _ := cmd.Output()
	return string(output)
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
	jsonOutput := run()
	json.Unmarshal([]byte(jsonOutput), &displays)
	return SetFrom(filterActive(displays))
}
