package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

func timerloop() {
	if len(os.Args) < 2 || os.Args[1] == "" {
		log.Println("no time stated")
		os.Exit(1)
	}

	input := os.Args[1]

	duration, err := time.ParseDuration(input)
	if err != nil {
		log.Fatalf("invalid duration %q: %v", input, err)
	}

	togglePause := make(chan struct{})
	go parseInput(togglePause)

	remaining := duration
	paused := false

	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	for remaining > 0 {
		select {
		case <-togglePause:
			paused = !paused
			if paused {
				fmt.Printf("\r%02d:%02d [PAUSED] type 'p' + Enter to resume\n", int(remaining.Minutes()), int(remaining.Seconds())%60)
			} else {
				fmt.Println("[RESUMED]")
				ticker.Reset(time.Second)
			}
		case <-ticker.C:
			if paused {
				continue
			}
			remaining -= time.Second
			printRemaining(remaining)
		}
	}

	fmt.Println("\r00:00 time elapsed    ")
}

func printRemaining(remaining time.Duration) {
	fmt.Printf("\r%02d:%02d remaining", int(remaining.Minutes()), int(remaining.Seconds())%60)
}

func parseInput(togglePause chan<- struct{}) {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		cmd := strings.TrimSpace(scanner.Text())
		if cmd == "p" {
			togglePause <- struct{}{}
		}
	}
}

func main() {
	timerloop()
}
