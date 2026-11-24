package main

import (
	"flag"
)

func main() {
	var pomodoro bool
	flag.BoolVar(&pomodoro, "p", false, "set the mode for the cli to pomodoro")
	flag.Parse()
	if !pomodoro {
		timer()
	} else {
		pomodorotimer()
	}
}
