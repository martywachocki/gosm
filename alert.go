package main

import "fmt"

// SendAlerts Sends the alerts for a services current status
func SendAlerts(service Service) {
	fmt.Println(service.Name + " is now " + service.Status)
}
