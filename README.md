# gosm (Work in Progress)
A program written in Golang for monitoring services using various protocols listed below. Alerts can be sent via SMTP and/or SMS (Twilio API).

### Supported Protocols
* HTTP
* HTTPS
* ICMP
* TCP

### Config
The application expects the configuration file to be named config.json. There is an example config file in the repo named config.example.json to use as a reference. Below is a brief description of each configuration item. All items are required unless explicity stated.
* **verbose** - Whether or not to print information to the console
* **check_interval** - How often to check each service that is in an online state (seconds)
* **pending_offline_interval** - How often to check each service is in a pending or offline state (seconds)
* **max_concurrent_checks** - The maximum concurrent checks
* **icmp_timeout** - Timeout threshold for ICMP (milliseconds)
* **successful_http_status_codes** - Which HTTP/HTTPS status codes are considered successful. Any status code not listed will be considered a failure response
* **ignore_http_cert_errors** - Whether or not to ignore HTTPS cert errors
* **failed_check_threshold** - How many consecutive failed checks are needed to consider a service offline
* **send_email** - Whether or not to send alerts via email
* **email_recipients** - Recipients of email alerts
* **smtp_host** - The SMTP server host to send emails from
* **smtp_port** - The SMTP server port
* **smtp_email_address** - The email address to send from
* **smtp_username** - The username for the SMTP server
* **smtp_password** - The password for the SMTP server
* **send_sms** - Whether or not to send alerts via sms
* **sms_recipients** - Recipients of sms alerts
* **twilio_account_sid** - Your Twilio Account SID
* **twilio_auth_token** - Your Twilio Auth Token
* **twilio_from_number** - Your Twilio phone number to send the SMS alerts from
* **services** - a list of services to monitor
    * **name** - The name of the service
    * **protocol** - The protocol of the service
    * **host** - The hostname or ip address of the service
    * **port** - The port of the service (optional for HTTP, HTTPS, and ICMP)


### TODO
* Implement SMTP and SMTP-TLS checks
* Optional limits email/sms alerts per second
* Create web UI
* Redesign to use sqlite database instead of JSON config files