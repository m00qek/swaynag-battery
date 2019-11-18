# swaynag-battery

Shows a message on top of all displays (using 
[swaynag](https://github.com/swaywm/sway/tree/master/swaynag)) when battery is 
discharging and bellow an appointed threshold. Also, closes the message when a
power supply is plugged in.

## Installing

Clone this repo, run `make build` and copy `bin/swaynag-battery` to 
somewhere in your `PATH`.

## API
```
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
```

## Acknowledgments

A special thanks to [Egor Kovetskiy](https://github.com/kovetskiy) who created
[i3-battery-nagbar](https://github.com/kovetskiy/i3-battery-nagbar), the program
that inspired me to create this one.
