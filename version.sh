latest_tag=$(git describe `git rev-list --tags --max-count=1`)

cat << EOF > version.go
package main

//go:generate bash ./version.sh
var version = "$latest_tag"
EOF
