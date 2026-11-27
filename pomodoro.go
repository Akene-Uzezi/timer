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
		p = "25m"
		b = "5m"
	} else {
		p = args[0]
		b = args[1]
	}

	pomodoroduration, err := time.ParseDuration(p)
	if err != nil {
		log.Fatalf("Invalid duration %q: %v", p, err)
	}

	breakduration, err := time.ParseDuration(b)
	if err != nil {
		log.Fatalf("Invalid duration %q: %v", b, err)
	}

	remaining := pomodoroduration
	paused := false
	breaktime := false
	cycle := 0

	togglePause := make(chan struct{})
	go parseInput(togglePause)

	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	for remaining > 0 {
		select {
		case <-togglePause:
			paused = !paused
			if paused {
				fmt.Printf("\033[2K\r%02d:%02d [PAUSED] type 'p' + Enter to resume", int(remaining.Minutes()), int(remaining.Seconds())%60)
			} else {
				fmt.Print("\033[2K\r[RESUMED]")
				ticker.Reset(time.Second)
			}
		case <-ticker.C:
			if paused {
				continue
			}
			remaining -= time.Second
			printRemainingPomodoro(remaining, cycle, breaktime)
			if remaining <= 0 {
				if !breaktime {
					cycle++
					if cycle == 4 {
						fmt.Print("\033[2K\rPomodoro done! Starting long break...\a")
						remaining = 15 * time.Minute
						cycle = 0
					} else {
						fmt.Print("\033[2K\rPomodoro done! Starting break...\a")
						remaining = breakduration
					}
				} else {
					fmt.Print("\033[2K\rBreak done! Starting next pomodoro\a")
					remaining = pomodoroduration
				}
				breaktime = !breaktime
			}
		}
	}
}
