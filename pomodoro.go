package main

import (
	"flag"
	"fmt"
	"log"
	"time"
)

// timer -p pomodorotime breaktime

func pomodorotimer() {
	args := flag.Args()
	var p string
	var b string
	if len(args) < 2 {
		log.Println("pomodoro cycle duration and break duration not stated. using default")
		p = "10s"
		b = "5m"
	} else {
		p = args[0]
		b = args[1]
	}

	pomodoroduration, err := time.ParseDuration(p)
	if err != nil {
		log.Fatalf("Invalid duration %q: %v", p, err)
	}

	_, err = time.ParseDuration(b)
	if err != nil {
		log.Fatalf("Invalid duration %q: %v", b, err)
	}

	remaining := pomodoroduration
	paused := false
	_ = false

	togglePause := make(chan struct{})
	go parseInput(togglePause)

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

	fmt.Println("\rdone           \a")
}
