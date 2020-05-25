# swaynag-battery

[![Release](https://img.shields.io/github/release/m00qek/swaynag-battery.svg?style=for-the-badge)](https://github.com/m00qek/swaynag-battery/releases/latest)
[![Software License](https://img.shields.io/badge/license-MIT-brightgreen.svg?style=for-the-badge)](/LICENSE.md)

Shows a message on top of all displays (using 
[swaynag](https://github.com/swaywm/sway/tree/master/swaynag)) when battery is 
discharging and bellow an appointed threshold. Also, closes the message when a
power supply is plugged in.

## Installing

Download the appropriate latest version binary from 
[releases](https://github.com/m00qek/swaynag-battery/releases) or clone this
repo, run `make build` and copy `bin/swaynag-battery` to somewhere in your
`PATH`.

In order to automatically run `swaynag-battery` when you execute Sway, you can 
configure [Sway to start using systemd](https://github.com/swaywm/sway/wiki/Systemd-integration)
and add a new user service file in 
`~/.config/systemd/user/swaynag-battery.service` with

```
[Unit]
Description=Low battery notification
PartOf=graphical-session.target

[Service]
Type=simple
ExecStart=/absolute/path/to/swaynag-battery

[Install]
WantedBy=sway-session.target
```

and then

```bash
systemctl --user enable swaynag-battery.service
systemctl --user start swaynag-battery.service
```

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

## Release Signatures

Releases are signed with 
[B7704FFB48AC73A1](https://keys.openpgp.org/vks/v1/by-fingerprint/2FC9D934AC901B875CAD71AAB7704FFB48AC73A1)
and published [on GitHub](https://github.com/m00qek/swaynag-battery/releases).

## Acknowledgments

A special thanks to [Egor Kovetskiy](https://github.com/kovetskiy) who created
[i3-battery-nagbar](https://github.com/kovetskiy/i3-battery-nagbar), the program
that inspired me to create this one.
