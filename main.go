package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"
)

func timerloop(togglePause chan bool) {
	if len(os.Args) < 2 || os.Args[1] == "" {
		log.Println("no time stated")
		os.Exit(1)
	}

	input := os.Args[1]

	duration, err := time.ParseDuration(input)
	if err != nil {
		log.Fatalf("invalid duration %q: %v", input, err)
	}

	for remaining := duration; remaining > 0; remaining -= time.Second {
		minutes := int(remaining.Minutes())
		seconds := int(remaining.Seconds()) % 60
		fmt.Printf("\r%02d:%02d remaining", minutes, seconds)
		time.Sleep(time.Second)

	}
	fmt.Println("\r00:00 time elapsed")
}

func parseInput(togglePause chan bool) {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		cmd := scanner.Text()
		if cmd == " " {
			togglePause <- true
		}
	}
}

func main() {
	togglePause := make(chan bool)
	go parseInput(togglePause)
	paused := false
	for {
		select {
		case <-togglePause:
			paused = !paused
			if paused {
				fmt.Println("\n[PAUSED] Type ' ' + Enter to resume...")
			} else {
				fmt.Println("[RESUMED]")
			}
		default:
			if !paused {
				timerloop(togglePause)
			} else {
				time.Sleep(100 * time.Millisecond)
			}
		}
	}
}
