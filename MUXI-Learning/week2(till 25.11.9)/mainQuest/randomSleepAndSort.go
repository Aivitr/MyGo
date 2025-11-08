package main

import (
	"fmt"
	"math/rand"
	"sort"
	"sync"
	"time"
)

type Data struct {
	randNum int
	ID      int
}

func main() {
	chData := make(chan Data, 20)
	Datas := make([]Data, 20)
	var wg sync.WaitGroup

	newRoutine := func(idx int) {
		randomNum := 1 + rand.Intn(1000)
		time.Sleep(time.Duration(randomNum) * time.Millisecond)
		chData <- Data{randomNum, idx}
		wg.Done()
	}

	for i := 1; i <= 20; i++ {
		wg.Add(1)
		go newRoutine(i)
	}

	for i := 0; i < 20; i++ {
		Datas[i] = <-chData
	}

	fmt.Println("未排序的数据为:")
	for i := range Datas {
		fmt.Printf("%d (with id :%d)\n", Datas[i].randNum, Datas[i].ID)
	}

	sort.Slice(Datas, func(i, j int) bool { return Datas[i].ID < Datas[j].ID })

	fmt.Println("按生成顺序排序后：")
	for i := range Datas {
		fmt.Printf("%d (with id :%d)\n", Datas[i].randNum, Datas[i].ID)
	}
}
