package notifier

import (
	"net/http"
	"time"
)

// CheckService checks if a service is reachable.
func CheckService(url string) bool {
	client := http.Client{
		Timeout: 10 * time.Second,
	}
	resp, err := client.Get(url)
	if err != nil || resp.StatusCode != http.StatusOK {
		return false
	}
	return true
}
