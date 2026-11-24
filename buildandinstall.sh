set -e

go build -o timer main.go

sudo mv timer /usr/bin/timer
