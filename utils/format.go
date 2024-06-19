package utils

import (
	"time"
)

func ConvertToLocalTime(unixTime int64, utcOffsetHours int) string {
	utcTime := time.Unix(unixTime, 0).UTC()
	localTime := utcTime.Add(time.Duration(utcOffsetHours) * time.Hour)
	return localTime.Format("2006-01-02 03:04 PM")
}

func GetWindDirection(degrees float64) string {
	directions := [...]string{"N", "NNE", "NE", "ENE", "E", "ESE", "SE", "SSE", "S", "SSW", "SW", "WSW", "W", "WNW", "NW", "NNW"}
	index := int((degrees+11.25)/22.5) % 16
	return directions[index]
}
