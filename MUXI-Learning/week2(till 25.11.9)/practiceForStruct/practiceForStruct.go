package main

import (
	"fmt"
	"time"
)

type employee struct {
	name   string
	salary float64
}
type myTime struct {
	time.Time
}

func main() {
	ben := employee{"ben", 3000.0}
	fmt.Println(ben.giveRaise(0.24))

	m := myTime{time.Now()}

	fmt.Println("Full time now:", m.String())
	fmt.Println("First 10 chars:", m.first10Chars())
}

func (e *employee) giveRaise(r float64) float64 {
	return e.salary * (1 + r)
}	//涨薪

func (t *myTime) first10Chars() string {
	return t.Time.String()[0:10]
}
