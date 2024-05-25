package model

import "time"

type PhoneNumberListResponse struct {
	PhoneNumbers []PhoneNumberResponse `json:"phone_numbers"`
}

type PhoneNumberResponse struct {
	Msisdn     string     `json:"msisdn"`
	Grade      string     `json:"grade"`
	Type       string     `json:"type"`
	ReservedAt *time.Time `json:"reserved_at"`
	ExpiresAt  *time.Time `json:"expires_at"`
}
