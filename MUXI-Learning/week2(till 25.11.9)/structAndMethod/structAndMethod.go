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

type Car struct {
	wheelCount int
}

type Mercedes struct {
	Car
}

type Log struct {
	msg string
}

type Customer struct {
	Name string
	log  *Log
}

func main() {
	/**********************************/

	ben := employee{"ben", 3000.0}

	fmt.Println(ben.giveRaise(0.24))

	/**********************************/

	m := myTime{time.Now()}

	fmt.Println("Full time now:", m.String())
	fmt.Println("First 10 chars:", m.first10Chars())

	/**********************************/

	mer := Mercedes{Car{4}}

	fmt.Println(mer.numberOfWheel())
	mer.sarHiToMerkel()

	/**********************************/

	c := new(Customer)
	c.Name = "Barak Obama"
	c.log = new(Log)
	c.Log().msg = "1 - Yes we can!" //c.log.msg
	c = &Customer{"Barak Obama", &Log{"1 - Yes we can!"}}
	c.Log().Add("2 - After me the world will be a better place!") //c.log.add
	fmt.Println(c.Log())                                          //fmt.Println(c.log)

	/**********************************/
}

/**********************************/
func (e *employee) giveRaise(r float64) float64 {
	return e.salary * (1 + r)
}

/**********************************/
func (t *myTime) first10Chars() string {
	return t.Time.String()[0:10]
}

/**********************************/
func (c *Car) numberOfWheel() int {
	return c.wheelCount
}

func (m *Mercedes) sarHiToMerkel() {
	fmt.Println("Hi,Merkel!")
}

/**********************************/
func (l *Log) Add(s string) {
	l.msg += "\n" + s
}

func (l *Log) String() string {
	return l.msg
}

func (c *Customer) Log() *Log {
	return c.log
}
