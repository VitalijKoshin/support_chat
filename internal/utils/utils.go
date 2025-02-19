package utils

import "time"

func ParseStrToTime(str string) (time.Time, error) {
	return time.Parse(time.RFC3339, str)
}
