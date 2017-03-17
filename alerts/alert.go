package alerts

import (
	"fmt"

	"../models"
)

// SendAlerts Sends the alerts for a services current status
func SendAlerts(service models.Service) {
	if models.CurrentConfig.Verbose {
		fmt.Println(service.Name + " is now " + service.Status)
	}
	if models.CurrentConfig.SendEmail {
		sendSMTPAlert(service)
	}
	if models.CurrentConfig.SendSMS {
		sendSMSAlert(service)
	}
}
