package main

import "io/ioutil"
import "encoding/json"

// Config The application configuration and settings
type Config struct {
	SuccessfulHTTPStatusCodes []int    `json:"successful_http_status_codes"`
	FailedCheckThreshold      int      `json:"failed_check_threshold"`
	SendEmail                 bool     `json:"send_email"`
	EmailRecipients           []string `json:"email_recipients"`
	SendSMS                   bool     `json:"send_sms"`
	SMSRecipients             []string `json:"sms_recipients"`
}

// ParseConfigFile Parses the config.json file
func ParseConfigFile() Config {
	file, err := ioutil.ReadFile("config.json")
	if err != nil {
		panic(err)
	}
	var config Config
	err = json.Unmarshal(file, &config)
	if err != nil {
		panic(err)
	}
	return config
}
