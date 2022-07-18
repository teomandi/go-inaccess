package utils

import (
	"fmt"
	"time"
)

// format of the given timestamps
const DateFormat = "20060102T150405Z"

// array with the accepted periods
var acceptedPeriods = [4]string{"1h", "1d", "1mo", "1y"}

// it parses a DateTime string according to the required format and it returns it as a time object
func ParseFormatedDate(dateTimeStr string) (time.Time, error) {
	t, err := time.Parse(DateFormat, dateTimeStr)
	if err != nil {
		fmt.Println(err)
		return time.Time{}, err
	}
	return t, nil
}

// it checks if the given period is accepted
func ValidatePeriod(p string) bool {
	for _, v := range acceptedPeriods {
		if v == p {
			return true
		}
	}
	return false
}
