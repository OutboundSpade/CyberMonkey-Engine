package engine

type DependencyInfo struct {
	Name      string
	Version   string
	Installed bool
}

type Dependency interface {
	Install() error   // Install the dependency (this should be idempotent)
	Uninstall() error // Uninstall the dependency (this should be idempotent)
	Info()
}

var dependencies = []Dependency{
	(*python)(nil),
}

func install_dependencies() error {
	// Install all dependencies
	for _, dep := range dependencies {
		err := dep.Install()
		if err != nil {
			return err
		}
	}
	return nil
}
