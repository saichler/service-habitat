package test

import (
	. "github.com/saichler/console/golang/console/commands"
	. "github.com/saichler/service-manager/golang/service-manager"
	. "github.com/saichler/utils/golang"
	"time"
)

type Test struct {
	running bool
	service IService
}

func NewTest(sm IService) *Test {
	sd := &Test{}
	sd.service = sm
	return sd
}

func (cmd *Test) Command() string {
	return "test"
}

func (cmd *Test) Description() string {
	return "Test Async"
}

func (cmd *Test) Usage() string {
	return "test"
}

func (cmd *Test) ConsoleId() *ConsoleId {
	return cmd.service.ConsoleId()
}

func (cmd *Test) RunCommand(args []string, id *ConsoleId) (string, *ConsoleId) {
	cmd.running = true
	for cmd.running {
		Println("Testing 1..2..3")
		time.Sleep(time.Second * 2)
	}
	return "", nil
}

func (cmd *Test) Stop() {
	cmd.running = false
	Println("Stop signal was sent")
}
