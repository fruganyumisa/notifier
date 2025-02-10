package models

// SMSRequest represents an SMS notification request.
type SMSRequest struct {
	Phone   string `json:"phone"`
	Message string `json:"message"`
}
