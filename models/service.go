package models

import (
	"crypto/tls"
	"net"
	"net/http"
	"strconv"
	"time"

	"github.com/sparrc/go-ping"
)

const (
	//Online The service is online
	Online = "ONLINE"
	// Pending The service is potentially offline, and will be marked so after meeting the Config.FailedCheckThreshold
	Pending = "PENDING"
	// Offline The service is offline
	Offline = "OFFLINE"
)

// Service Represents a service that is being monitored
type Service struct {
	Name         string `json:"name"`
	Protocol     string `json:"protocol"`
	Host         string `json:"host"`
	Port         int    `json:"port"`
	Status       string `json:"status"`
	FailureCount int    `json:"failure_count"`
}

var (
	CurrentConfig Config
)

// CheckService Checks whether a service is online or offline
func (service *Service) CheckService() bool {
	switch service.Protocol {
	case "http", "https":
		return checkHTTP(service.Host)
	case "icmp":
		return checkICMP(service.Host)
	case "tcp":
		return checkTCP(service.Host, service.Port)
	default:
		panic("Unsupported protocol: " + service.Protocol)
	}
}

func checkHTTP(host string) bool {
	transport := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: CurrentConfig.IgnoreHTTPSCertErrors},
	}
	client := &http.Client{Transport: transport}
	response, err := client.Get(host)
	if err != nil {
		return false
	}
	defer response.Body.Close()
	var responseStatusCode = response.StatusCode
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
	connection, err := net.Dial("tcp", host+":"+strconv.Itoa(port))
	defer connection.Close()
	if err != nil {
		return false
	}
	return true
}
