package models

import "time"

type PingStatus struct {
	ID          int       `json:"id"`
	IPAddress   string    `json:"ip_address"`
	PingTime    time.Time `json:"ping_time"`
	Success     bool      `json:"success"`
	LastSuccess time.Time `json:"last_success"`
}
