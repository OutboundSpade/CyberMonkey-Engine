package engine

import (
	"fmt"
	"log"

	"github.com/takama/daemon"
)

func daemonize() error {
	service, err := daemon.New("name", "description", daemon.SystemDaemon)
	if err != nil {
		log.Fatal("Error: ", err)
	}
	status, err := service.Install()
	if err != nil {
		log.Fatal(status, "\nError: ", err)
	}
	fmt.Println(status)
	return nil
}
