package typeutil

import "time"

func Time(t time.Time) *time.Time {
	return &t
}

func String(s string) *string {
	return &s
}
