package model

import "time"

type PingStatus struct {
	IPAddress   string    `json:"ip_address"`
	PingTime    time.Time `json:"ping_time"`
	Success     bool      `json:"success"`
	LastSuccess time.Time `json:"last_success"`
}
