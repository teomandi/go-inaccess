package strcucts

import (
	"time"

	"github.com/teomandi/go-inaccess/utils"
)

type Task struct {
	Period   string
	Timezone string
	T1       time.Time
	T2       time.Time
}

func (t *Task) GetMatchingTimestamps() []string {
	timestamps := []string{}

	T1 := t.T1.Round(time.Hour)
	i := 0
	for {
		var timestamp time.Time
		if t.Period == "1h" {
			// timestamp = time.Date(T1.Year(), T1.Month(), T1.Day(), T1.Hour()+i, 0, 0, 0, T1.Location()) //hour
			timestamp = T1.Add(time.Hour * time.Duration(i))
		} else if t.Period == "1d" {
			// timestamp = time.Date(T1.Year(), T1.Month(), T1.Day()+i, T1.Hour(), 0, 0, 0, T1.Location()) //day
			timestamp = T1.AddDate(0, 0, i)
		} else if t.Period == "1mo" {
			// timestamp = time.Date(T1.Year(), T1.Month()+1+time.Month(i), 0, T1.Hour(), 0, 0, 0, T1.Location()) //month
			timestamp = T1.AddDate(0, i, 0)
			timestamp = time.Date(timestamp.Year(), timestamp.Month()+1, 0, timestamp.Hour(), 0, 0, 0, timestamp.Location())
		} else if t.Period == "1y" {
			// timestamp = time.Date(T1.Year()+1+i, T1.Month()-1, 0, T1.Hour(), 0, 0, 0, T1.Location()) //year
			timestamp = T1.AddDate(i+1, 0, 0)
			timestamp = time.Date(timestamp.Year(), timestamp.Month()-1, 0, timestamp.Hour(), 0, 0, 0, timestamp.Location())
		}

		if timestamp.After(t.T2) {
			break
		}
		timestamps = append(timestamps, timestamp.Format(utils.DateFormat))
		i++
	}
	return timestamps
}
