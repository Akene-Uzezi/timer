set -e

go build -o timer .

sudo mv timer /usr/bin/timer
