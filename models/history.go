package models

import "time"

type History struct {
	CustomerID string    `json:"customer_id"`
	Action     string    `json:"action"`
	Timestamp  time.Time `json:"timestamp"`
}
