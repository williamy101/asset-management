package util

import (
	"fmt"
	"time"
)

func ParseDate(dateStr string) (*time.Time, error) {
	parsed, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		return nil, fmt.Errorf("invalid date format: %s, expected YYYY-MM-DD", dateStr)
	}
	return &parsed, nil
}
