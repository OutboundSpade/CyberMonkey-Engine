package engine

import (
	"fmt"
	"log"
	"strings"

	"github.com/takama/daemon"
)

func daemonize() error {
	service, err := daemon.New("cybermonkey-engine", "cybermonkey scoring engine", daemon.SystemDaemon)
	if err != nil {
		log.Fatal("Error: ", err)
	}
	status, err := service.Status()
	if err == nil {
		if strings.Contains(status, "stopped") {
			status, err := service.Start()
			if err != nil {
				return fmt.Errorf(status, "\nError: ", err)
			}
		}
		return nil
	}

	status, err = service.Install("run")
	if err != nil {
		return fmt.Errorf(status, "\nError: ", err)
	}
	fmt.Println(status)
	status, err = service.Start()
	if err != nil {
		return fmt.Errorf(status, "\nError: ", err)
	}
	fmt.Println(status)

	return nil
}
