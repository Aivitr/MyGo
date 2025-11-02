package main

import "fmt"

//最小公倍数

func lcm(m, n int)int{
	a, b := m, n
	for b != 0 {
		a, b = b, a % b
	}
	return (m*n)/a
}


func main(){
	x, y, m, n, L := 0, 0, 0, 0, 0

	fmt.Scanln(&x,&y,&m,&n,&L)

	px, py := x, y
	pxs := []int{x}
	pys := []int{y}

	t := 1

	//男生可能所在位置
	for {

		px = (px + m)%L

		if px != pxs[0] {
			pxs = append(pxs, px)
			t++
		} else {
			t = 1	
			break
		}

	} 

	//女生可能所在位置
	for {
		py = (py + n)%L

		if py != pys[0] {
			pys = append(pys, py)
			t++
		} else {	
			t = 1 
			break
		}

	} 


	//遍历一个公周期
	attempt := 0 
	max_attempt := lcm(len(pxs),len(pys))

	for {
		if attempt > max_attempt {
			fmt.Println("-1")
			break
		}

		if pxs[(attempt)%len(pxs)] == pys[(attempt)%len(pys)] {
			fmt.Println(attempt)
			break
		} else {
			attempt ++
			continue
		}
	} 
	
}