package main

import (
	"fmt"

	"github.com/sfreiberg/gotwilio"
)

// SendAlerts Sends the alerts for a services current status
func SendAlerts(service Service) {
	if CurrentConfig.Verbose {
		fmt.Println(service.Name + " is now " + service.Status)
	}
	if CurrentConfig.SendEmail {
		sendEmailAlert(service)
	}
	if CurrentConfig.SendSMS {
		sendSMSAlert(service)
	}
}

func sendEmailAlert(service Service) {

}

func sendSMSAlert(service Service) {
	twilio := gotwilio.NewTwilioClient(CurrentConfig.TwilioAccountSID, CurrentConfig.TwilioAuthToken)
	for _, number := range CurrentConfig.SMSRecipients {
		_, exception, err := twilio.SendSMS(CurrentConfig.TwilioPhoneNumber, number, "gosm - "+service.Name+" ("+service.Protocol+") is now "+service.Status, "", "")
		if exception != nil {
			fmt.Println(err)
		}
		if err != nil {
			fmt.Println(err)
		}
	}
}
