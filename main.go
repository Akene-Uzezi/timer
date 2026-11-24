package main

import (
	"flag"
	"fmt"
)

func main() {
	var pomodoro bool
	flag.BoolVar(&pomodoro, "p", false, "set the mode for the cli to pomodoro")
	flag.Parse()
	if !pomodoro {
		timer()
	} else {
		fmt.Printf("pomodoro: %v", pomodoro)
	}
}
