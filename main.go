package main

import (
	"time"
)

type Parameters struct {
	displays  StringSet
	interval  time.Duration
	message   string
	threshold int
	uevent    string
}

func DesiredDisplays(displays StringSet, activeDisplays StringSet) StringSet {
	if len(displays) == 0 {
		return activeDisplays
	}
	return Intersection(displays, activeDisplays)
}

func tick(watcher *Watcher, params Parameters) {
	battery, _ := LoadBatteryInfo(params.uevent)
	displays := DesiredDisplays(params.displays, ActiveDisplays())

	if !battery.Charging() && battery.Capacity <= params.threshold {
		messages := ShowAll(params.message, watcher.MessagesFor(displays))
		watcher.Update(messages, battery.Status)
	}

	if battery.Charging() && battery.Status != watcher.Status() {
		messages := watcher.Messages()
		CloseAll(messages)
		watcher.Update(messages, battery.Status)
		watcher.CleanUp(displays)
	}
}

func main() {
	interval, _ := time.ParseDuration("1s")

	params := Parameters{
		displays:  SetFrom([]string{"eDP-1"}),
		interval:  interval,
		message:   "You battery is running low. Please plug in a power adapter",
		threshold: 100,
		uevent:    "/sys/class/power_supply/BAT0/uevent"}

	watcher := NewWatcher()

	tick(&watcher, params)
	for range time.Tick(params.interval) {
		tick(&watcher, params)
	}
}
