package model

import "time"

type Activity struct {
	Model
	IP           string        `json:"ip" form:"ip" `
	Method       string        `json:"method" form:"method" `
	Path         string        `json:"path" form:"path" `
	Status       int           `json:"status" form:"status" `
	Latency      time.Duration `json:"latency" form:"latency" `
	Agent        string        `json:"agent" form:"agent" `
	ErrorMessage string        `json:"error_message" form:"error_message" `
	Body         string        `json:"body" form:"body" `
	Resp         string        `json:"resp" form:"resp"`
}
