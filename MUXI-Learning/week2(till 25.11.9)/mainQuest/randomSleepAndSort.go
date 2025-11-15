package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Data struct {
	randNum int
	ID      int
}

func main() {
	chData := make(chan Data, 20)
	done := make(chan bool)

	newRoutine := func(idx int) {
		randomNum := 1 + rand.Intn(1000)
		time.Sleep(time.Duration(randomNum) * time.Millisecond)
		chData <- Data{randomNum, idx}
		done <- true
	}

	for i := 1; i <= 20; i++ {
		go newRoutine(i)
	}

	for i := 1; i <= 20; i++ {
		<-done
	}
	close(done)

	fmt.Println("按生成顺序输出")

	cnt := 0
	for v := range chData {
		if cnt == 20 {break}
		cnt++

		fmt.Println(v.randNum, "with id :", v.ID)
		chData <- v
	}

	fmt.Println("按id输出")
	for i := 1; i <= 20; i++ {
		remain := 20 - i
		found := false 
		
		for j := 0; j <= remain && !found; j++ {
			select {
			case v := <-chData:
				if v.ID == i {
					fmt.Println(v.randNum, "with id :", v.ID)
					found = true 
				} else {
					chData <- v 
				}
			}
		}
		
	}
}