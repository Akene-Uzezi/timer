package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
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
