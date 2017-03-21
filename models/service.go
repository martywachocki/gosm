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
	ID           int           `db:"id" json:"id"`
	Name         string        `db:"name" json:"name"`
	Protocol     string        `db:"protocol" json:"protocol"`
	Host         string        `db:"host" json:"host"`
	Port         jsonNullInt64 `db:"port" json:"port"`
	Status       string        `json:"status"`
	FailureCount int           `json:"failure_count"`
}

var (
	// CurrentServices The currently monitored services
	CurrentServices []Service
)

// CheckService Checks whether a service is online or offline
func (service *Service) CheckService() bool {
	switch service.Protocol {
	case "http", "https":
		return checkHTTP(service.Host)
	case "icmp":
		return checkICMP(service.Host)
	case "tcp":
		return checkTCP(service.Host, service.Port.Int64)
	default:
		panic("Unsupported protocol: " + service.Protocol)
	}
}

func checkHTTP(host string) bool {
	transport := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: CurrentConfig.IgnoreHTTPSCertErrors},
	}
	client := &http.Client{
		Transport: transport,
		Timeout:   time.Millisecond * time.Duration(CurrentConfig.ConnectionTimeout),
	}
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
	pinger.Timeout = time.Millisecond * time.Duration(CurrentConfig.ConnectionTimeout)
	pinger.Run()
	statistics := pinger.Statistics()
	return statistics.PacketsSent == statistics.PacketsRecv
}

func checkTCP(host string, port int64) bool {
	dialer := &net.Dialer{
		Timeout: time.Millisecond * time.Duration(CurrentConfig.ConnectionTimeout),
	}
	connection, err := dialer.Dial("tcp", host+":"+strconv.FormatInt(port, 10))
	if err != nil {
		return false
	}
	defer connection.Close()
	return true
}

// LoadServices Loads all the services into CurrentServices and sets defaults
func LoadServices() {
	var services []Service
	Database.Select(&services, "SELECT * FROM services")

	for i := range services {
		services[i].Status = Online
	}
	CurrentServices = services
}
