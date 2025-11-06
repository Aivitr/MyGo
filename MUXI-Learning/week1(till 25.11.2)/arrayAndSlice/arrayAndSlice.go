package main

import (
	"fmt"
	"unicode/utf8"
)

//尝试修改切片元素
func changeElement(s []int) {
	s[0] = 999
}

//追加元素后修改切片元素
func appendAndChange(s []int) {
	s = append(s, 6)
	s[0] = 999
}

func main() {

	arr1 := [5]int{1,2,3,4,5}
	arr2 := [5]int{1,2,3,4,5}

	s1 := arr1[1:3]
	s2 := arr2[3:5]//cap(s) = len(s),已满

	changeElement(s1)// 这里传递进去的是切片副本，但是副本的指针还是指向arr

    fmt.Println(arr1)//  [1 999 3 4 5]（可以看到底层数组被修改)
	fmt.Println(s1)//  [999 3]（底层数组arr被改变，那么切片元素自然也就改变）
	fmt.Println(s1) //   [2 3]

 	appendAndChange(s2)//本次扩容导致切片副本指针指向新的匿名数组

    fmt.Println(s2) //   [4 5]（切片元素完全没变)
    fmt.Println(arr2)//  [1,2,3,4,5]（原本的底层数组也没变） 

    str1, str2 := "hello  world", "the"
	inserstr := InsertStringSlice(str1,str2,6)
	fmt.Println(inserstr)

	str3 := "你好,Golang"
	fmt.Println(len(str3),utf8.RuneCountInString(str3))


}

func InsertStringSlice(s1 string, s2 string, pos int)(string){
	if pos < 0 {pos = 0}

	if pos > len(s1) {pos = len(s1)}

	return s1[:pos] + s2 + s1[pos:]
}

