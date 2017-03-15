package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// CurrentConfig The current application configuration
var (
	CurrentConfig     Config
	checkChannel      = make(chan *Service)
	checkCountChannel chan (bool)
)

func main() {
	CurrentConfig = ParseConfigFile()
	checkCountChannel = make(chan bool, CurrentConfig.MaxConcurrentChecks)

	// Workaround for SIGTERM not working when pinging
	c := make(chan os.Signal, 2)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		os.Exit(1)
	}()

	go processChecks()
	go checkOnlineServices()
	checkPendingOfflineServices()
}

func checkOnlineServices() {
	for {
		for i := range CurrentConfig.Services {
			if CurrentConfig.Services[i].Status != Online {
				checkCountChannel <- true
				checkChannel <- &CurrentConfig.Services[i]
			}
		}
		time.Sleep(time.Second * time.Duration(CurrentConfig.PendingOfflineCheckInterval))
	}
}

func checkPendingOfflineServices() {
	for {
		for i := range CurrentConfig.Services {
			if CurrentConfig.Services[i].Status == Online {
				checkCountChannel <- true
				checkChannel <- &CurrentConfig.Services[i]
			}
		}
		time.Sleep(time.Second * time.Duration(CurrentConfig.CheckInterval))
	}
}

func processChecks() {
	for {
		service := <-checkChannel
		online := service.CheckService()
		if online == true {
			if service.Status == Offline {
				service.Status = Online
				go SendAlerts(*service)
			} else if service.Status == Pending {
				service.Status = Online
			}
			service.FailureCount = 0
		} else {
			if service.Status == Online {
				service.Status = Pending
				service.FailureCount = 1
				if CurrentConfig.Verbose {
					fmt.Println(service.Name + " is now in the " + service.Status + " state")
				}
			} else if service.Status == Pending {
				service.FailureCount++
				if service.FailureCount >= CurrentConfig.FailedCheckThreshold {
					service.Status = Offline
					go SendAlerts(*service)
				}
			}
		}
		<-checkCountChannel
	}
}
