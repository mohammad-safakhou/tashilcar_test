package models

import "time"

type RegisterEndpoint struct {
	Endpoint   EndpointRules `json:"endpoint"`
	Scheduling Scheduling    `json:"scheduling"`
}

type EndpointRules struct {
	Url          string              `json:"url"`
	Method       string              `json:"method"`
	Header       map[string][]string `json:"header"`
	Body         string              `json:"body"`
	EndpointLogs []EndpointLog       `json:"endpoint_logs"`
}

type Scheduling struct {
	Repeater int       `json:"repeater"` // 1m, 5m, 15m, 1h, 4h, 24h
	ExpireAt time.Time `json:"expire_at"`
}

type EndpointLog struct {
	Status  bool      `json:"status"`
	Created time.Time `json:"created"`
}
