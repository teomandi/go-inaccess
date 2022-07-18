package structs

import (
	"time"

	"github.com/teomandi/go-inaccess/pkg/utils"
)

type Task struct {
	Period   string
	Timezone string
	T1       time.Time
	T2       time.Time
}

// it creates and returns a Task after it validates the given arguments.
func MakeTask(p string, tz string, t1Str string, t2Str string) (Task, bool, string) {
	time.LoadLocation(tz)
	t1, err := utils.ParseFormatedDate(t1Str)
	if err != nil {
		return Task{}, false, "Invalid format at t1: " + t1Str
	}
	t2, err := utils.ParseFormatedDate(t2Str)
	if err != nil {
		return Task{}, false, "Invalid format at t2: " + t2Str
	}
	if !utils.ValidatePeriod(p) {
		return Task{}, false, "Unsupported period"
	}
	return Task{Period: p, Timezone: "Europe/Athens", T1: t1, T2: t2}, true, ""
}

// Returns the matching timestamps of the task
func (t *Task) GetMatchingTimestamps() []string {
	timestamps := []string{}

	T1 := t.T1.Round(time.Hour)
	i := 0
	for {
		var timestamp time.Time
		if t.Period == "1h" {
			timestamp = T1.Add(time.Hour * time.Duration(i))
		} else if t.Period == "1d" {
			timestamp = T1.AddDate(0, 0, i)
		} else if t.Period == "1mo" {
			timestamp = T1.AddDate(0, i, 0)
			timestamp = time.Date(timestamp.Year(), timestamp.Month()+1, 0, timestamp.Hour(), 0, 0, 0, timestamp.UTC().Location())
		} else if t.Period == "1y" {
			timestamp = T1.AddDate(i+1, 0, 0)
			timestamp = time.Date(timestamp.Year(), timestamp.Month()-1, 0, timestamp.Hour(), 0, 0, 0, timestamp.UTC().Location())
		}
		if timestamp.After(t.T2) {
			break
		}
		timestamps = append(timestamps, timestamp.Format(utils.DateFormat))
		i++
	}
	return timestamps
}
