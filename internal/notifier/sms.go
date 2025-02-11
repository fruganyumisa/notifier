package notifier

import (
	"io"
	"net/http"
	"net/url"
	"strings"
	"sync"
)

// SendSMS sends an SMS notification to multiple admins via the SMS gateway.
func SendSMS(gatewayURL string, phones []string, message string, senderHeader string) error {
	var wg sync.WaitGroup
	var err error

	for _, phone := range phones {
		wg.Add(1)
		go func(phone string) {
			defer wg.Done()

			// Create form data
			formData := url.Values{}
			formData.Set("dst", phone)    // Add phone number
			formData.Set("text", message) // Add message
			formData.Set("src", senderHeader)
			formData.Set("register", "final")

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
			defer func(Body io.ReadCloser) {
				err := Body.Close()
				if err != nil {

				}
			}(resp.Body)

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
