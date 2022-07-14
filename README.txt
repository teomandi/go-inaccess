https://gist.github.com/hosszukalman/8fd4bb046f4a20d55f4723776d6cc964
https://www.techieindoor.com/go-add-years-months-days-to-current-date-in-go/



	t1Str := "20180214T204603Z"
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

	t1 = t1.Round(time.Hour)
	i := 1
	for {
		// timestamp := time.Date(t1.Year(), t1.Month(), t1.Day()+i, t1.Hour(), 0, 0, 0, t1.Location()) //day
		// timestamp := time.Date(t1.Year(), t1.Month()+time.Month(i), 0, t1.Hour(), 0, 0, 0, t1.Location()) //month
		timestamp := time.Date(t1.Year()+i, t1.Month()-1, 0, t1.Hour(), 0, 0, 0, t1.Location()) //year
		if timestamp.After(t2) {
			break
		}
		fmt.Println(timestamp)
		i++
	}

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