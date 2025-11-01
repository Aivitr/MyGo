package main

import (
	"fmt"
)

func main(){
	testStr :=  "hello,世界"
	for idx := range testStr { fmt.Print(idx, " ")}
	//共循环8次，依次输出0 1 2 3 4 5 6 9
}
