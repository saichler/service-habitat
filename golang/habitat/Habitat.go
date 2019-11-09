package main

import (
	commands2 "github.com/saichler/file-management-service/golang/file-management-service/commands"
	message_handlers2 "github.com/saichler/file-management-service/golang/file-management-service/message-handlers"
	service2 "github.com/saichler/file-management-service/golang/file-management-service/service"
	"github.com/saichler/service-habitat/golang/habitat/test"
	"github.com/saichler/service-management-service/golang/management-service/commands"
	message_handlers "github.com/saichler/service-management-service/golang/management-service/message-handlers"
	"github.com/saichler/service-management-service/golang/management-service/service"
	"github.com/saichler/service-manager/golang/service-manager"
	"io/ioutil"
	"os"
	"strings"
)

func main() {

	serviceManager, e := service_manager.NewServiceManager()
	if e != nil {
		return
	}

	serviceManager.Console().RegisterCommand(commands.NewLS(serviceManager), "ls")
	serviceManager.Console().RegisterCommand(commands.NewCD(serviceManager), "cd")
	serviceManager.AddService(NewFileService())
	ms, mh, mc := NewService()
	serviceManager.AddService(ms, mh, mc)

	serviceManager.Console().RegisterCommand(test.NewTest(ms), "test")

	files, e := ioutil.ReadDir("./plugins")
	if e == nil {
		for _, f := range files {
			if strings.Contains(f.Name(), ".so") {
				serviceManager.LoadService("./plugins/" + f.Name())
			}
		}
	} else {
		os.Create("./plugins")
	}
	serviceManager.WaitForShutdown()
}

func NewService() (service_manager.IService, service_manager.IServiceCommands, service_manager.IServiceMessageHandlers) {
	s := &service.ManagementService{}
	h := &message_handlers.ManagementHandlers{}
	h.Init(s)
	c := &commands.ManagementCommands{}
	c.Init(s, h)
	return s, c, h
}

func NewFileService() (service_manager.IService, service_manager.IServiceCommands, service_manager.IServiceMessageHandlers) {
	s := &service2.FileManagerService{}
	h := &message_handlers2.FileManagerHandlers{}
	h.Init(s)
	c := &commands2.FileManagerCommands{}
	c.Init(s, h)
	return s, c, h
}
