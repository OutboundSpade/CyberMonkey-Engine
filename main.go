package main

import "github.com/OutboundSpade/CyberMonkey-Engine/engine"

type BundleConfig struct {
	Modules []string `yaml:"modules"`
}

func main() {
	must(engine.Start())
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
