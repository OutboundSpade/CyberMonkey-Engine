package engine

import (
	"fmt"
	"os"
	"os/user"
	"runtime"
)

func start() error {
	err := ensure_is_admin()
	if err != nil {
		return err
	}
	switch runtime.GOOS {
	case "windows":
		return fmt.Errorf("This is an unsupported OS!")
	case "linux":
		err := setup_linux()
		if err != nil {
			return err
		}
	default:
		return fmt.Errorf("This is an unsupported OS!")
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
