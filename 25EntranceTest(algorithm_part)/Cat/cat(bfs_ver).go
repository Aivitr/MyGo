package main

import (
	"fmt"
)


type State struct {
	pos    int   //当前位置
	energy int   //当前能量上限
	steps  int   //步数
}

func main(){
	n, E , r := 0, 0, 0
	var road string

	fmt.Scanln(&n,&E,&r)
	fmt.Scanln(&road)

	ans := -1 

	//特殊的，有极端情况：终点或起点为-
	if road[0] == '-' || road[n] == '-' {

		fmt.Println(ans)

	} else {

	//一般的，采用BFS,切片模拟队列
	//map嵌套,应对E极大时的内存消耗问题

	visited := make([]map[int]bool, n)

	for i := range visited {visited[i] = make(map[int]bool)}
	
	
	queue := []State{{pos: 0, energy: E, steps: 0}}
	visited[0][E] = true


	found := false


	for len(queue) > 0 {

		//出列
		cur := queue[0]
		queue = queue[1:]

		//遍历所有落点
		for k := 1; k <= cur.energy; k++ {

			//更新位置
			newPos := cur.pos + k

			//是否到达终点
			if newPos == n-1 {
				ans = cur.steps + 1
				found = true
				break
			}

			//是否超出道路范围
			if newPos >= n {
				continue
			}

			//是否已访问 
			if visited[newPos][cur.energy] {
				continue
			}

			//是否为陷阱
			if road[newPos] == '-' {
				continue
			}

			//更新E值
			newEnergy := cur.energy
			if road[newPos] == '+' {
				newEnergy += r
			}

			//标记当前位置pos为已访问,入列
			visited[newPos][newEnergy] = true
			queue = append(queue, State{
				pos:newPos,
				energy:newEnergy,
				steps:cur.steps + 1,
			})
		}		
		if found {break}
	}

	fmt.Println(ans)

    }
}