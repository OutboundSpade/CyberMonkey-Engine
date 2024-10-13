package engine

import (
	"fmt"
	"os"
	"os/user"
	"runtime"
)

const (
	INSTALL_DIR_LINUX   = "/opt/cybermonkey-engine"
	INSTALL_DIR_WINDOWS = "C:\\Program Files\\CyberMonkey-Engine"
)

func Start() error {
	err := ensure_is_admin()
	if err != nil {
		return err
	}

	// err = daemonize()
	// if err != nil {
	// 	return err
	// }

	err = install_dependencies()
	if err != nil {
		return err
	}

	return nil
}

func ensure_is_admin() error {
	switch runtime.GOOS {
	case "linux":
		current_user, err := user.Current()
		if err != nil {
			return err
		}
		if current_user.Uid != "0" {
			return fmt.Errorf("You must run this program as root!")
		}
	case "windows":
		_, err := os.Open("\\\\.\\PHYSICALDRIVE0")
		if err != nil {
			return fmt.Errorf("You must run this program as an administrator!")
		}
	default:
		return fmt.Errorf("This is an unsupported OS!")
	}
	return nil
}
