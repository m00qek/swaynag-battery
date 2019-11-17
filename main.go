package main

import (
	"time"
)

type BatteryWatcher struct {
	Messages       []MessageProcess
	PreviousStatus Status
}

type Parameters struct {
	displays  []string
	interval  time.Duration
	message   string
	threshold int
	uevent    string
}

func tick(watcher *BatteryWatcher, params Parameters) {
	battery, _ := LoadBatteryInfo(params.uevent)
	statusHasChanged := battery.Status != watcher.PreviousStatus

	if battery.Status != CHARGING && battery.Capacity <= params.threshold {
		watcher.Messages = ShowAll(params.message, watcher.Messages)
	} else if battery.Status == CHARGING && statusHasChanged {
		CloseAll(watcher.Messages)
	}

	watcher.PreviousStatus = battery.Status
}

func createMessages(displays []string) []MessageProcess {
	return []MessageProcess{
		MessageProcess{PID: 0, Display: "eDP-1"},
		MessageProcess{PID: 0, Display: "HDMI-A-1"},
	}
}

func main() {
	interval, _ := time.ParseDuration("5m")

	params := Parameters{
		displays:  []string{},
		interval:  interval,
		message:   "You battery is running low. Please plug in a power adapter",
		threshold: 15,
		uevent:    "/sys/class/power_supply/BAT0/uevent"}

	watcher := &BatteryWatcher{Messages: createMessages(params.displays)}

	tick(watcher, params)
	for range time.Tick(params.interval) {
		tick(watcher, params)
	}
}
