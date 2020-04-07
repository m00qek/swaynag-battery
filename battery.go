package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type Status uint8

const (
	UNKNOWN      Status = 0
	CHARGING     Status = 1
	DISCHARGING  Status = 2
	NOT_CHARGING Status = 3
	FULL         Status = 4
)

type Battery struct {
	Name       string
	ModelName  string
	Technology string
	Capacity   int
	Status     Status
}

func (battery *Battery) Charging() bool {
	return battery.Status == FULL || battery.Status == CHARGING
}

func parseStatus(value string) Status {
	switch value {
	case "Charging":
		return CHARGING
	case "Discharging":
		return DISCHARGING
	case "Not charging":
		return NOT_CHARGING
	case "Full":
		return FULL
	default:
		return UNKNOWN
	}
}

func load(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func parse(content []string) map[string]string {
	info := make(map[string]string)

	for _, line := range content {
		tokens := strings.SplitN(line, "=", 2)
		if len(tokens) != 2 {
			continue
		}

		key, value := tokens[0], tokens[1]
		info[key] = value
	}

	return info
}

func build(info map[string]string) (*Battery, error) {
	capacity, err := strconv.Atoi(info["POWER_SUPPLY_CAPACITY"])
	if err != nil {
		return nil, err
	}

	return &Battery{
		Name:       info["POWER_SUPPLY_NAME"],
		Capacity:   capacity,
		ModelName:  info["POWER_SUPPLY_MODEL_NAME"],
		Status:     parseStatus(info["POWER_SUPPLY_STATUS"]),
		Technology: info["POWER_SUPPLY_TECHNOLOGY"],
	}, nil
}

func LoadBatteryInfo(uevent string) (*Battery, error) {
	content, err := load(uevent)
	if err != nil {
		logError("ERROR: Could not load battery file '%s'.", uevent)
		return nil, err
	}

	battery, err := build(parse(content))
	if err != nil {
		logError("Could not parse 'POWER_SUPPLY_CAPACITY' from battery file '%s'.", uevent)
		return nil, err
	}

	return battery, nil
}
