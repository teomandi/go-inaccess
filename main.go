package main

import (
	"fmt"

	strcucts "github.com/teomandi/go-inaccess/structs"
	"github.com/teomandi/go-inaccess/utils"
)

func main() {
	t1Str := "20180214T204603Z"
	t2Str := "20211115T123456Z"
	period := "1y"

	t1, err := utils.ParseFormatedDate(t1Str)
	if err != nil {
		fmt.Println(err)
		return
	}
	t2, err := utils.ParseFormatedDate(t2Str)
	if err != nil {
		fmt.Println(err)
		return
	}

	task := strcucts.Task{Period: period, Timezone: "Europe/Athens", T1: t1, T2: t2}
	timestamps := task.GetMatchingTimestamps()

	for _, v := range timestamps {
		fmt.Printf("%v\n", v)
	}

}
