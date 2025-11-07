package main

import "fmt"

func newCnter() func() int {

    cnt := 0

    return func() int {
        cnt++ 
        return cnt 
    }
}

func main() {
 
    cnter1 := newCnter()
    fmt.Println("计数器1第1次调用：", cnter1()) // 输出：1
    fmt.Println("计数器1第2次调用：", cnter1()) // 输出：2
    fmt.Println("计数器1第3次调用：", cnter1()) // 输出：3

    fmt.Println("\n")

    cnter2 := newCnter()
    fmt.Println("计数器2第1次调用：",  cnter2()) // 输出：1
    fmt.Println("计数器2第2次调用：",  cnter2()) // 输出：2
    fmt.Println("计数器1第4次调用：",  cnter1()) // 输出：4
}