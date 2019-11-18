package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	docopt "github.com/docopt/docopt-go"
)

type Parameters struct {
	displays  StringSet
	interval  time.Duration
	message   string
	threshold int
	uevent    string
}

var (
	usage = `
Shows a message (using swaynag) when battery percentage is less then specified
value.

Usage:
  swaynag-battery [options]
  swaynag-battery -h | --help
  swaynag-battery --version

Options:
  --displays <display-list>  Comma separated list of displays to show the
                             alert - the default is to show in all displays.
  --threshold <int>          Percentual threshold to show notification.
                             [default: 15]
  --interval <duration>      Check battery at every interval. [default: 5m]
  --uevent <path>            Uevent path for reading battery stats.
                             [default: /sys/class/power_supply/BAT0/uevent]
  -h --help                  Show this screen.
  --version                  Show version.
`
)

func CommandLineParameters() Parameters {
	args, err := docopt.ParseDoc(usage)
	if err != nil {
		fmt.Println("err", err)
		panic(err)
	}

	interval, _ := time.ParseDuration(args["--interval"].(string))
	threshold, _ := strconv.Atoi(args["--threshold"].(string))

	displays, ok := args["--displays"].(string)
	if !ok {
		displays = ""
	}

	return Parameters{
		displays:  SetFrom(strings.Split(displays, ",")),
		interval:  interval,
		message:   "You battery is running low. Please plug in a power adapter",
		threshold: threshold,
		uevent:    args["--uevent"].(string)}
}
