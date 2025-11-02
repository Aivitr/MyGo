package main

import (
	"fmt"
	"sort"
)

func main(){

	k,num := 0, 0
	goods,sum := []int{},[]int{-1}

	//处理输入
	fmt.Scan(&k)
	for i := 0; i < k; i++ {
		fmt.Scan(&num)
		goods = append(goods, num)
	}

	//筛选指定商品
	for i := 1; i < k-2; i++  {
		if goods[i-1]%2 == 0 && goods[i]%2 == 1 && goods[i+1]%2== 0 {
			sum = append(sum, goods[i-1] + goods[i] + goods[i+1])
		} 
	}

	//排序 
	sort.Slice(sum, func(i, j int)bool{return sum[i] > sum[j]})
	
	//输出
	fmt.Println(sum[0])
}