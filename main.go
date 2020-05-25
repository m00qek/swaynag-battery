package main

import (
	"os"
	"time"
)

func DesiredDisplays(displays StringSet, activeDisplays StringSet) StringSet {
	if len(displays) == 0 {
		return activeDisplays
	}

	return Intersection(displays, activeDisplays)
}

func tick(watcher *Watcher, params Parameters) {
	battery, err := LoadBatteryInfo(params.uevent)
	if err != nil {
		logWarning("Skipping this cycle due to errors occurred.")
		return
	}

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
	params := CommandLineParameters(os.Args[1:])
	watcher := NewWatcher()

	tick(&watcher, params)
	for range time.Tick(params.interval) {
		tick(&watcher, params)
	}
}
