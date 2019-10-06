package utils

import "time"

func RoundToDay(t time.Time) time.Time {
	return t.UTC().Truncate(24 * time.Hour)
}
