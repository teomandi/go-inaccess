package structs

type TestTask struct {
	Period   string
	Timezone string
	T1       string
	T2       string
	Result   bool
}

// returns the test-task as a string
func (t *TestTask) ToString() string {
	return t.Period + " " + t.Timezone + " " + t.T1 + "-" + t.T2
}
