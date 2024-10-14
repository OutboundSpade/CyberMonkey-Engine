package engine

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"runtime"

	"gopkg.in/yaml.v3"
)

type Script struct {
	Script string  `yaml:"script"`
	Type   *string `yaml:"type"`
}

func (s *Script) UnmarshalYAML(node *yaml.Node) error {
	// Unmarshal the script
	scriptMap := map[string]string{}
	err := node.Decode(&scriptMap)
	if err != nil {
		return err
	}
	s.Script = scriptMap["script"]

	if t, ok := scriptMap["type"]; ok {
		s.Type = &t
	} else {
		ext := path.Ext(s.Script)
		if t, ok := extensions[ext]; ok {
			s.Type = &t
		} else {
			s.Type = &unknown
		}
	}
	return nil
}

type Module struct {
	Path        string  `yaml:"-"`
	Name        string  `yaml:"name"`
	Description string  `yaml:"description"`
	Points      *int    `yaml:"points"`
	Break       *Script `yaml:"break"`
	Check       Script  `yaml:"check"`
}

var unknown = "unknown"
var extensions = map[string]string{
	".py":  "python",
	".sh":  "bash",
	".ps1": "powershell",
	".ps":  "powershell",
	".bat": "batch",
}

func getModules() ([]Module, error) {
	modulesPath := ""
	switch runtime.GOOS {
	case "linux":
		modulesPath = path.Join(INSTALL_DIR_LINUX, "modules")
	case "windows":
		modulesPath = path.Join(INSTALL_DIR_WINDOWS, "modules")
	default:
		return nil, fmt.Errorf("Unsupported OS: %s", runtime.GOOS)
	}

	modules := []Module{}
	// Traverse the modules directory and parse the mod.yaml files
	err := filepath.Walk(modulesPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() || (info.Name() != "mod.yaml" && info.Name() != "mod.yml") {
			return nil
		}
		if info.Name() == "mod.yaml" || info.Name() == "mod.yml" {
			// Parse the mod.yaml file
			fmt.Printf("Parsing %s\n", path)
			mod := Module{}
			file, err := os.ReadFile(path)
			if err != nil {
				return err
			}
			err = yaml.Unmarshal(file, &mod)
			if err != nil {
				return err
			}
			mod.Path = filepath.Dir(path)
			modules = append(modules, mod)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return modules, nil
}
