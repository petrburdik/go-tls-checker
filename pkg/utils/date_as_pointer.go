package utils

import (
	"fmt"
	"time"
)

// GetDatePtr - get date as pointer
func GetDatePtr(year int, month time.Month, day int, hour int, min int, sec int, nsec int, loc *time.Location) *time.Time {
	date := time.Date(year, month, day, hour, min, sec, nsec, loc)
	return &date
}

// GetUnixDatePtr - get date, GMT timezone, milliseconds since unix epoch
func GetUnixDatePtr(nsec int64) *time.Time {
	date := time.UnixMilli(nsec)
	return &date
}

func ParseDatePtr(layout string, value string) *time.Time {
	date, err := time.Parse(layout, value)
	if err != nil {
		fmt.Printf("error parse date %q", err.Error())
		return nil
	}
	return &date
}
