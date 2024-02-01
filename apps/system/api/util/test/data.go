package test

import "time"

func GetDefaultTime() time.Time {
	defaultTime, _ := time.Parse("2006-01-02 15:04:05", "2024-01-10 12:00:00")
	return defaultTime
}
