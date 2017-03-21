package models

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

// Config The application configuration and settings
type Config struct {
	Verbose                     bool     `json:"verbose"`
	WebUIHost                   string   `json:"web_ui_host"`
	WebUIPort                   int      `json:"web_ui_port"`
	CheckInterval               int      `json:"check_interval"`
	PendingOfflineCheckInterval int      `json:"pending_offline_check_interval"`
	MaxConcurrentChecks         int      `json:"max_concurrent_checks"`
	ConnectionTimeout           int      `json:"connection_timeout"`
	SuccessfulHTTPStatusCodes   []int    `json:"successful_http_status_codes"`
	IgnoreHTTPSCertErrors       bool     `json:"ignore_https_cert_errors"`
	FailedCheckThreshold        int      `json:"failed_check_threshold"`
	SendEmail                   bool     `json:"send_email"`
	EmailRecipients             []string `json:"email_recipients"`
	SMTPHost                    string   `json:"smtp_host"`
	SMTPPort                    int      `json:"smtp_port"`
	SMTPEmailAddress            string   `json:"smtp_email_address"`
	SMTPUsername                string   `json:"smtp_username"`
	SMTPPassword                string   `json:"smtp_password"`
	SendSMS                     bool     `json:"send_sms"`
	SMSRecipients               []string `json:"sms_recipients"`
	TwilioAccountSID            string   `json:"twilio_account_sid"`
	TwilioAuthToken             string   `json:"twilio_auth_token"`
	TwilioPhoneNumber           string   `json:"twilio_phone_number"`
}

var (
	// CurrentConfig The current configuration
	CurrentConfig Config
)

// ParseConfigFile Parses the config.json file
func ParseConfigFile() Config {
	if len(os.Args) < 2 {
		panic("Expected run syntax: './gosm /path/to/config.json'")
	}
	file, err := ioutil.ReadFile(os.Args[1])
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
