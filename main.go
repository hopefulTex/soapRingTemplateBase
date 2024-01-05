package main

import (
	"log"
	"soapRingTest/ui"
)

func main() {
	// handle args (or not)

	// config depends on args
	p := ui.NewProgram()
	_, err := p.Run()
	if err != nil {
		log.Fatalf("\nerror: %s\n", err.Error())
	}
}
