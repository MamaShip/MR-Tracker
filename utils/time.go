package utils

import "time"

func ParseTime(time_str string) (t time.Time, err error) {
	t, err = time.Parse(time.RFC3339Nano, time_str)
	if err != nil {
		return t, err
	}
	return t, nil
}
