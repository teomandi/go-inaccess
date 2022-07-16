package utils

import (
	"fmt"
	"time"
)

const DateFormat = "20060102T150405Z"

var acceptedPeriods = [4]string{"1h", "1d", "1mo", "1y"}

func ParseFormatedDate(dateTimeStr string) (time.Time, error) {
	t, err := time.Parse(DateFormat, dateTimeStr)
	if err != nil {
		fmt.Println(err)
		return time.Time{}, err
	}
	return t, nil
}

func ValidatePeriod(p string) bool {
	for _, v := range acceptedPeriods {
		if v == p {
			return true
		}
	}
	return false
}
