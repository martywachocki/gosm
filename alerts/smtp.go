package alerts

import (
	"fmt"
	"net/smtp"
	"strconv"

	"../models"
)

func sendSMTPAlert(service models.Service) {
	auth := smtp.PlainAuth("",
		models.CurrentConfig.SMTPUsername,
		models.CurrentConfig.SMTPPassword,
		models.CurrentConfig.SMTPHost)

	message := "Subject: [gosm] " + service.Name + " is " + service.Status + "\r\n"
	message += service.Name + " is now " + service.Status + "\r\n"
	message += "Protocol: " + service.Protocol + "\r\n"
	message += "Host: " + service.Host + "\r\n"
	if service.Port.Value != nil {
		message += "Port: " + strconv.FormatInt(service.Port.Int64, 10) + "\r\n"
	}
	err := smtp.SendMail(
		models.CurrentConfig.SMTPHost+":"+strconv.Itoa(models.CurrentConfig.SMTPPort),
		auth, models.CurrentConfig.SMTPEmailAddress,
		models.CurrentConfig.EmailRecipients,
		[]byte(message))
	if err != nil {
		fmt.Println(err)
	}
}
