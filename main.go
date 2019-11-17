package main

import (
	"fmt"
)

func main() {
	fmt.Println(LoadBatteryInfo("/sys/class/power_supply/BAT0/uevent"))
}
