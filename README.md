# gosm (Work In Progress)
A program written in Golang for monitoring services using various protocols listed below.

### Supported Protocols
* HTTP
* HTTPS
* ICMP
* SMTP
* SMTP-TLS
* TCP

### Config
There is an example config file in the repo named config.example.json to use as a start. Below is a brief description of each configuration item.
* **check_interval** - How often to check each service that is in an online state (seconds)
* **pending_offline_interval** - How often to check each service is in a pending or offline state (seconds)
* **max_concurrent_checks** - The maximum concurrent checks
* **icmp_timeout** - The ICMP timeout in milliseconds
* **successful_http_status_codes** - Which HTTP/HTTPS status codes are considered successful. Any status code not listed will be considered a failure response
* **ignore_http_cert_errors** - Whether or not to ignore HTTPS cert errors
* **failed_check_threshold** - How many consecutive failed checks are needed to consider a service offline
* **send_email** - Whether or not to send alerts via email
* **email_recipients** - Recipients of email alerts
* **send_sms** - Whether or not to send alerts via sms
* **sms_recipients** - Recipients of sms alerts
* **services** - a list of services to monitor
    * **name** - The name of the service
    * **protocol** - The protocol of the service
    * **host** - The hostname or ip address of the service
    * **port** - The port of the service