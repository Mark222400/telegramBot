package main

import "time"

func date() (arrOfTime []string) {
	x := time.Now()
	first := time.Now().Format("02.01.2006")
	arrOfTime = []string{first}
	for i := 0; i < 27; i++ {
		x = x.Add(24 * time.Hour)
		el := x.Format("02.01.2006")
		arrOfTime = append(arrOfTime, el)
	}
	return
}



