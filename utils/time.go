package utils

import "time"

func ParseTime(time_str string) (t time.Time, err error) {
	if time_str == "" {
		return time.Time{}, nil
	}
	t, err = time.Parse(time.RFC3339Nano, time_str)
	if err != nil {
		return t, err
	}
	return t, nil
}
