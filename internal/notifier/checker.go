package notifier

import (
	"log"
	"net/http"
	"time"
)

// CheckService checks if a service is reachable.
func CheckService(url string) bool {
	client := http.Client{
		Timeout: 10 * time.Second,
	}
	resp, err := client.Get(url)
	if err == nil && resp.StatusCode >= 200 && resp.StatusCode <= 599 {

		return true
	}
	log.Printf("Service Down Error Response received from service: %v", err)

	return false
}
