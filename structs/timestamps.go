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
	i := 1
	for {
		var timestamp time.Time
		if t.Period == "1h" {
			timestamp = time.Date(T1.Year(), T1.Month()-1, 0, T1.Hour()+i, 0, 0, 0, T1.Location()) //year
		} else if t.Period == "1d" {
			timestamp = time.Date(T1.Year(), T1.Month(), T1.Day()+i, T1.Hour(), 0, 0, 0, T1.Location()) //day
		} else if t.Period == "1m" {
			timestamp = time.Date(T1.Year(), T1.Month()+time.Month(i), 0, T1.Hour(), 0, 0, 0, T1.Location()) //month
		} else if t.Period == "1y" {
			timestamp = time.Date(T1.Year()+i, T1.Month()-1, 0, T1.Hour(), 0, 0, 0, T1.Location()) //year
		}

		if timestamp.After(t.T2) {
			break
		}
		timestamps = append(timestamps, timestamp.Format(utils.DateFormat))
		i++
	}
	return timestamps
}
