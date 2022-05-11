package main

import (
	"os"
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
  --message <message>        Use a custom low battery message. 
                             [default: "Your battery is running low. Please 
			     plug in a power adapter"]
  -h --help                  Show this screen.
  --version                  Show version.
`
)

func CommandLineParameters(arguments []string) Parameters {
	args, err := docopt.ParseArgs(usage, arguments, version)
	if err != nil {
		logAndExit(18, "Unable to parse input arguments.")
	}

	interval, err := time.ParseDuration(args["--interval"].(string))
	if err != nil {
		logAndExit(28, "Unable to parse '--interval %s': the value must be a duration.", args["--interval"])
	}

	threshold, err := strconv.Atoi(args["--threshold"].(string))
	if err != nil {
		logAndExit(38, "Unable to parse '--threshold %s': the value must be an integer number.", args["--threshold"])
	}

	displays := []string{}
	d, ok := args["--displays"].(string)
	if ok {
		displays = strings.Split(d, ",")
	}

	uevent := args["--uevent"].(string)
	file, err := os.Open(uevent)
	if err != nil {
		logAndExit(42, "Could not load battery file '%s'.", uevent)
	}
	file.Close()

	return Parameters{
		displays:  SetFrom(displays),
		interval:  interval,
		message:   message,
		threshold: threshold,
		uevent:    uevent}
}
