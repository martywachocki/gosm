package main

import (
	"encoding/json"
	"io/ioutil"
)

// Config The application configuration and settings
type Config struct {
	CheckInterval               int       `json:"check_interval"`
	PendingOfflineCheckInterval int       `json:"pending_offline_check_interval"`
	MaxConcurrentChecks         int       `json:"max_concurrent_checks"`
	ICMPTimeout                 int       `json:"icmp_timeout"`
	SuccessfulHTTPStatusCodes   []int     `json:"successful_http_status_codes"`
	IgnoreHTTPCertErrors        bool      `json:"ignore_http_cert_errors"`
	FailedCheckThreshold        int       `json:"failed_check_threshold"`
	SendEmail                   bool      `json:"send_email"`
	EmailRecipients             []string  `json:"email_recipients"`
	SendSMS                     bool      `json:"send_sms"`
	SMSRecipients               []string  `json:"sms_recipients"`
	Services                    []Service `json:"services"`
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
	config.setupServices()
	return config
}

func (config *Config) setupServices() {
	for i, _ := range config.Services {
		config.Services[i].Status = Online
		config.Services[i].FailureCount = 0
	}
}
