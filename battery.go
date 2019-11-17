package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type Battery struct {
	Name       string
	ModelName  string
	Technology string
	Capacity   int
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
		Technology: info["POWER_SUPPLY_TECHNOLOGY"],
	}, nil
}

func LoadBatteryInfo(uevent string) (*Battery, error) {
	content, err := load(uevent)
	if err != nil {
		return nil, err
	}

	return build(parse(content))
}
