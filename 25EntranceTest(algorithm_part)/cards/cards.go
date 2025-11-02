package main

import (
	"fmt"
	"sort"
)

func main(){

	n := 0
	fmt.Scan(&n)

	cardsA := make([]int, n)
	cardsB := make([]int, n)

	for i := range cardsA {fmt.Scan(&cardsA[i])}
	for i := range cardsB {fmt.Scan(&cardsB[i])}


	//贪心策略

	sort.Ints(cardsA)
	sort.Ints(cardsB)

	aL, aR := 0, n-1
	bL, bR := 0, n-1
	golds := 0

	for aL <= aR || bL <= bR {

		if cardsA[aL] > cardsB[bL] {
			golds += 200
			aL ++
			bL ++
		} else if cardsA[aR] > cardsB[bR] {
			golds += 200
			aR --
			bR --
		} else {
			if cardsA[aL] < cardsB[bR] {golds -= 200}
			aL ++
			bR --
		}
	}

	fmt.Println(golds)
}