package model

import (
	"time"
)

type PhoneNumber struct {
	Msisdn     string
	Grade      string
	Type       string
	ReservedAt *time.Time
	ExpiresAt  *time.Time
}
