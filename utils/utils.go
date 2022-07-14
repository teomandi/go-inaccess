package utils

import (
	"fmt"
	"time"
)

const DateFormat = "20060102T150405Z"

func ParseFormatedDate(dateTimeStr string) (time.Time, error) {
	t, err := time.Parse(DateFormat, dateTimeStr)
	if err != nil {
		fmt.Println(err)
		return time.Time{}, err
	}
	return t, nil
}
