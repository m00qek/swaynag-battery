package main

import (
	"encoding/json"
)

type Display struct {
	Name   string
	Active bool
}

func run() (string, error) {
	response, err := sendIpc(swayGetOutputs)

	if err != nil {
		logError("Unable to get sway outputs.")
		return "", err
	}

	return response, nil
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
