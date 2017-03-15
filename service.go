package main

import (
	"crypto/tls"
	"net/http"
	"time"

	ping "github.com/sparrc/go-ping"
)

const (
	//Online The service is online
	Online = "Online"
	// Pending The service is potentially offline, and will be marked so after meeting the Config.FailedCheckThreshold
	Pending = "Pending"
	// Offline The service is offline
	Offline = "Offline"
)

// Service Represents a service that is being monitored
type Service struct {
	Name         string `json:"name"`
	Protocol     string `json:"protocol"`
	Host         string `json:"host"`
	Port         int    `json:"port"`
	Status       string
	FailureCount int
}

// CheckService Checks whether a service is online or offline
func (service *Service) CheckService() bool {
	switch service.Protocol {
	case "http", "https":
		return checkHTTP(service.Host, service.Port)
	case "icmp":
		return checkICMP(service.Host)
	case "tcp":
		return checkTCP(service.Host, service.Port)
	case "smtp":
		return checkSMTP(service.Host, service.Port, false)
	case "smtp-tls":
		return checkSMTP(service.Host, service.Port, true)
	default:
		panic("Unsupported protocol: " + service.Protocol)
	}
}

func checkHTTP(host string, port int) bool {
	transport := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: CurrentConfig.IgnoreHTTPCertErrors},
	}
	client := &http.Client{Transport: transport}
	resp, err := client.Get(host)
	if err != nil {
		return false
	}
	var responseStatusCode = resp.StatusCode
	var isValidStatusCode = false
	for _, statusCode := range CurrentConfig.SuccessfulHTTPStatusCodes {
		if responseStatusCode == statusCode {
			isValidStatusCode = true
			break
		}
	}
	return isValidStatusCode
}

func checkICMP(host string) bool {
	pinger, err := ping.NewPinger(host)
	if err != nil {
		return false
	}
	pinger.Count = 1
	pinger.Timeout = time.Millisecond * time.Duration(CurrentConfig.ICMPTimeout)
	pinger.Run()
	statistics := pinger.Statistics()
	return statistics.PacketsSent == statistics.PacketsRecv
}

func checkTCP(host string, port int) bool {
	return false
}

func checkSMTP(host string, port int, tls bool) bool {
	return false
}
