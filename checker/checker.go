package checker

import (
	"fmt"
	"time"

	"../alerts"
	"../models"
)

var (
	checkChannel      = make(chan *models.Service)
	checkCountChannel chan (bool)
)

// Start Starts the service checker process
func Start() {
	checkCountChannel = make(chan bool, models.CurrentConfig.MaxConcurrentChecks)
	go processChecks()
	go checkOnlineServices()
	checkPendingOfflineServices()
}

func checkOnlineServices() {
	for {
		for i := range models.CurrentServices {
			if len(models.CurrentServices) <= i {
				break
			}
			if models.CurrentServices[i].Status != models.Online {
				checkCountChannel <- true
				checkChannel <- &models.CurrentServices[i]
			}
		}
		time.Sleep(time.Second * time.Duration(models.CurrentConfig.PendingOfflineCheckInterval))
	}
}

func checkPendingOfflineServices() {
	for {
		for i := range models.CurrentServices {
			if len(models.CurrentServices) <= i {
				break
			}
			if models.CurrentServices[i].Status == models.Online {
				checkCountChannel <- true
				checkChannel <- &models.CurrentServices[i]
			}
		}
		time.Sleep(time.Second * time.Duration(models.CurrentConfig.CheckInterval))
	}
}

func processChecks() {
	for {
		service := <-checkChannel
		online := service.CheckService()
		if online == true {
			if service.Status == models.Offline {
				service.Status = models.Online
				go alerts.SendAlerts(*service)
			} else if service.Status == models.Pending {
				service.Status = models.Online
			}
			service.FailureCount = 0
		} else {
			if service.Status == models.Online {
				service.Status = models.Pending
				service.FailureCount = 1
				if models.CurrentConfig.Verbose {
					fmt.Println(service.Name + " is now in the " + service.Status + " state")
				}
			} else if service.Status == models.Pending {
				service.FailureCount++
				if service.FailureCount >= models.CurrentConfig.FailedCheckThreshold {
					service.Status = models.Offline
					go alerts.SendAlerts(*service)
				}
			}
		}
		<-checkCountChannel
	}
}
