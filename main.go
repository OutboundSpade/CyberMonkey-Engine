package main

import (
	"os"

	"github.com/OutboundSpade/CyberMonkey-Engine/engine"
)

func main() {
	if len(os.Args) > 1 && os.Args[1] == "run" {
		must(engine.RunBackground())
		return
	}
	must(engine.Start())
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
