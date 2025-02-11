# Notifier

The **Notifier** is a Go-based application that periodically checks the connectivity of remote services and sends SMS notifications to administrators if any service becomes unreachable. It is designed to be simple, reliable, and easy to configure.

---

## Features

- **Periodic Service Checks**: Monitors services at a configurable interval (e.g., every 5 minutes).
- **SMS Notifications**: Sends SMS alerts to multiple administrators if a service is down.
- **Graceful Shutdown**: Handles termination signals (e.g., `Ctrl+C`) gracefully.
- **Configurable**: Supports configuration via environment variables or a YAML file.
- **TCP Connectivity Checks**: Uses TCP (Telnet) to verify service availability.

---

## Getting Started

### Prerequisites

- Go 1.20 or higher
- An SMS gateway API endpoint (for sending notifications)

---

### Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/fruganyumisa/notifier.git
   cd notifier
   ```

2. Install dependencies:
   ```bash
   go mod tidy
   ```

3. Build the application:
   ```bash
   go build -o notifier ./cmd/notifier
   ```

---

## Configuration

The service can be configured using either a **Environment File** or **os environment variables**.

### YAML Configuration

Create a `config.env` file in the root directory:

```bash
SERVICES="host1:8080,host2:22"
CHECK_INTERVAL="5m"
SMS_GATEWAY_URL="http://sms-gateway.com/send
```

### Environment Variables

Alternatively, you can configure the service using environment variables:

```bash
export SERVICES="host1:8080,host2:22"
export CHECK_INTERVAL="5m"
export SMS_GATEWAY_URL="http://sms-gateway.com/api/send"
export ADMIN_PHONES="+1234567890,+0987654321"
```

---

## Running the Service

Start the service using the following command:

```bash
./notifier
```

---

## SMS Message Template

When a service is down, the following SMS message is sent to administrators:

```
🚨 Service Alert 🚨

Service: {ServiceName}
Status: Down
Time: {Timestamp}

Details:
{FailureDetails}

Action Required:
Please investigate immediately. Check server logs and restart if necessary.
```

---

## Graceful Shutdown

The service supports graceful shutdown. When a termination signal (e.g., `Ctrl+C`) is received, it stops the monitoring process and cleans up resources before exiting.

---

## Folder Structure

```
notifier/
├── cmd/
│   └── notifier/
│       └── main.go          # Entry point of the application
├── internal/
│   ├── config/
│   │   └── config.go        # Configuration management
│   ├── notifier/
│   │   ├── checker.go       # Service connectivity checker
│   │   ├── sms.go           # SMS notification logic
│   │   └── notifier.go      # Main notifier logic
│   ├── models/
│   │   └── models.go        # Data models (e.g., SMSRequest)
│   └── utils/
│       └── utils.go         # Utility functions (e.g., joinStrings)
├── config.yaml              # Configuration file (optional)
├── go.mod                   # Go module file
├── go.sum                   # Go dependencies checksum file
└── README.md                # Project documentation
```

---

## Example Logs

### Service Start
```
Starting notifier service...
Notifier started. Waiting for the first tick...
```

### Service Check
```
Tick received. Checking services...
Checking service: host1:8080
Connection to host1:8080 succeeded.
Service host1:8080 is up.
Checking service: host2:22
Connection to host2:22 failed: dial tcp ...
Service host2:22 is down.
Sending SMS notification: Services down: host2:22
SMS notification sent successfully.
```

### Graceful Shutdown
```
^CReceived signal: interrupt. Shutting down...
Received shutdown signal. Stopping notifier...
Notifier service stopped.
```

---

## Contributing

Contributions are welcome! Please follow these steps:

1. Fork the repository.
2. Create a new branch for your feature or bugfix.
3. Submit a pull request with a detailed description of your changes.

---

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

---

## Support

For questions or issues, please open an issue on the [GitHub repository](https://github.com/fruganyumisa/notifier).

---

## Acknowledgments

- Built with ❤️ using Go.
- Inspired by the need for reliable service monitoring and alerting.

---

Enjoy using the  **Notifier**! 