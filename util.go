package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

func printRemaining(remaining time.Duration) {
	fmt.Printf("\033[2K\r%02d:%02d remaining", int(remaining.Minutes()), int(remaining.Seconds())%60)
}

func parseInput(togglePause chan<- struct{}) {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		cmd := strings.TrimSpace(scanner.Text())
		if cmd == "p" {
			togglePause <- struct{}{}
		}
		if cmd == "s" {
			os.Exit(0)
		}
	}
}
