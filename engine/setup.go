package engine

import "os"

func setup_linux() error {
	const engine_path = "/opt/cybermonkey-engine"

	// Make directory if it doesn't exist
	if _, err := os.Stat(engine_path); os.IsNotExist(err) {
		err := os.Mkdir(engine_path, 0755)
		if err != nil {
			return err
		}
	}

	// Change owner of directory
	err := os.Chown(engine_path, 0, 0)
	if err != nil {
		return err
	}

	return nil
}
