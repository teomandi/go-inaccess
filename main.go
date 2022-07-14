package main

import (
	"fmt"
	"time"
)

func main() {
	// DateTime Format: YYYYMMDDTHHmmSSZ
	dateFormat := "20060102T150405Z"

	t1Str := "20211010T204603Z"
	t1, err := time.Parse(dateFormat, t1Str)
	if err != nil {
		fmt.Println(err)
		return
	}
	t2Str := "20211115T123456Z"
	t2, err := time.Parse(dateFormat, t2Str)
	if err != nil {
		fmt.Println(err)
		return
	}

	//day
	i := 1
	for {
		timestamp := time.Date(t1.Year(), t1.Month(), t1.Day()+i, t1.Hour(), 0, 0, 0, t1.Location()) //end of month
		if timestamp.After(t2) {
			break
		}
		fmt.Println(timestamp)
		i++
	}

	//month
	// i := 1
	// for {
	// 	timestamp := time.Date(t1.Year(), t1.Month()+time.Month(i), 0, t1.Hour(), 0, 0, 0, t1.Location()) //end of month
	// 	if timestamp.After(t2) {
	// 		break
	// 	}
	// 	fmt.Println(timestamp)
	// 	i++
	// }

	//year
	// i := 1
	// for {
	// 	timestamp := time.Date(t1.Year()+i, t1.Month(), 0, t1.Hour()+1, 0, 0, 0, t1.Location()) //end of month
	// 	if timestamp.After(t2) {
	// 		break
	// 	}
	// 	fmt.Println(timestamp)
	// 	i++
	// }

	// timestamp = timestamp.AddDate(0, 1, 0) //end of month
	// timestamp = time.Date(timestamp.Year(), timestamp.Month()+1, 0, timestamp.Hour(), 0, 0, 0, timestamp.Location())
	// fmt.Println(timestamp)
	// for {

	// }
	//
	// fmt.Println(timestamp)
	// i++
	// timestamp2 := time.Date(timestamp.Year(), timestamp.Month()+time.Month(i), 0, 0, 0, 0, 0, timestamp.Location())
	// fmt.Println(timestamp2)

	// timestamp := time.Date(t1.Year(), t1.Month()+1, 0, 0, 0, 0, 0, t1.Location())
	// fmt.Println(timestamp)

	// timestamp = timestamp.AddDate(0, 1, 0)
	// timestamp = time.Date(timestamp.Year(), timestamp.Month()+1, 0, 0, 0, 0, 0, timestamp.Location())
	// fmt.Println(timestamp)

	// for {
	// 	i++
	// 	timestamp := time.Date(timestamp.Year(), timestamp.Month()+2, 0, 0, 0, 0, 0, timestamp.Location())
	// 	printDateTime(timestamp)
	// 	if timestamp.After(t2) {
	// 		break
	// 	}

	// }

}

func printDateTime(timestamp time.Time) {
	// fmt.Printf("%d) %d/%02d/%02d T %02d:%02d:%02d Z\n", i, timestamp.Year(), timestamp.Month(), timestamp.Day(), timestamp.Hour(), timestamp.Minute(), timestamp.Second())
	fmt.Println(timestamp.Format("20060102T150405Z"))
}
