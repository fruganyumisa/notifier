package notifier

import (
	"net/http"
	"net/url"
	"strings"
	"sync"
)

// SendSMS sends an SMS notification to multiple admins via the SMS gateway.
func SendSMS(gatewayURL string, phones []string, message string) error {
	var wg sync.WaitGroup
	var err error

	for _, phone := range phones {
		wg.Add(1)
		go func(phone string) {
			defer wg.Done()

			// Create form data
			formData := url.Values{}
			formData.Set("phone", phone)     // Add phone number
			formData.Set("message", message) // Add message

			// Create a new HTTP request
			req, reqErr := http.NewRequest("POST", gatewayURL, strings.NewReader(formData.Encode()))
			if reqErr != nil {
				err = reqErr
				return
			}

			// Set headers
			req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

			// Send the request
			client := &http.Client{}
			resp, httpErr := client.Do(req)
			if httpErr != nil {
				err = httpErr
				return
			}
			defer resp.Body.Close()

			// Check the response status
			if resp.StatusCode != http.StatusOK {
				err = httpErr
				return
			}
		}(phone)
	}

	wg.Wait()
	return err
}
