package alerts

import (
	"fmt"

	"github.com/sfreiberg/gotwilio"

	"github.com/martywachocki/gosm/models"
)

func sendSMSAlert(service models.Service) {
	twilio := gotwilio.NewTwilioClient(models.CurrentConfig.TwilioAccountSID, models.CurrentConfig.TwilioAuthToken)
	for _, number := range models.CurrentConfig.SMSRecipients {
		_, exception, err := twilio.SendSMS(models.CurrentConfig.TwilioPhoneNumber, number, "[gosm] "+service.Name+" ("+service.Protocol+") is now "+service.Status, "", "")
		if exception != nil {
			fmt.Println(err)
		}
		if err != nil {
			fmt.Println(err)
		}
	}
}
