package models

import "time"

type URL struct {
	URL					string			`json:"url"`
	Alias			   	string			`json:"short"`
	Expiry				time.Duration	`json:"expiry"`
	XRateRemaining		int				`json:"rate_limit"`	
	XRateLimitReset 	time.Duration	`json:"rate_limit_reset"`
}