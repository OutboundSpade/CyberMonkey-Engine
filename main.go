package main

import (
	"flag"
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type BundleConfig struct {
	Modules []string `yaml:"modules"`
}

func main() {
	bundle := flag.NewFlagSet("bundle", flag.ExitOnError)
	bundleConfigPath := bundle.String("config", "config.yml", "Path to the config file")
	bundleModulesPath := bundle.String("modules", "modules", "Path to the modules directory")
	bundleOutputPath := bundle.String("output", "cybermonkey-comp.zip", "Path to the output file")
	flagSets := []*flag.FlagSet{bundle}
	help := flag.Bool("h", false, "Show help")
	flag.Parse()
	if len(os.Args) < 2 || *help {
		for _, flagSet := range flagSets {
			fmt.Println(flagSet.Name())
			flagSet.PrintDefaults()
		}
		os.Exit(1)
	}
	switch os.Args[1] {
	case "bundle":
		bundle.Parse(os.Args[2:])
		config := BundleConfig{}
		// Read the config file
		file, err := os.Open(*bundleConfigPath)
		must(err)
		defer file.Close()
		must(yaml.NewDecoder(file).Decode(&config))
		must(createBundle(&config, *bundleModulesPath, *bundleOutputPath))
	}
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
