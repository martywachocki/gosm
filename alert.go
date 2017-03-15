package main

import (
	"fmt"
	"net/smtp"
	"strconv"

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
	auth := smtp.PlainAuth("",
		CurrentConfig.SMTPUsername,
		CurrentConfig.SMTPPassword,
		CurrentConfig.SMTPHost)

	message := "Subject: [gosm] " + service.Name + " is " + service.Status + "\r\n"
	message += service.Name + " is now " + service.Status + "\r\n"
	message += "Protocol: " + service.Protocol + "\r\n"
	message += "Host: " + service.Host + "\r\n"
	message += "Port: " + strconv.Itoa(service.Port) + "\r\n"

	err := smtp.SendMail(
		CurrentConfig.SMTPHost+":"+strconv.Itoa(CurrentConfig.SMTPPort),
		auth, CurrentConfig.SMTPEmailAddress,
		CurrentConfig.EmailRecipients,
		[]byte(message))
	if err != nil {
		fmt.Println(err)
	}
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
