package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
)

func main() {
	buf := bytes.NewBuffer([]byte("this is a test for bytes package"))
	// 分割字符串
	s := buf.String()
	var part1, part2 string
	if len(s) >= 10 {
		part1 = s[:10]
		part2 = s[10:]
	} else {
		part1 = s
		part2 = ""
	}
	fmt.Println("part1:", part1)
	fmt.Println("part2:", part2)

	buf.WriteString("let's try reset")
	buf.Reset()
	if buf.Len() == 0 {
		fmt.Println("(Here's nothing!)")
	}

	// 读取到 ":"
	buf.WriteString("try to get string before the colon : ?????????")
	str, err := buf.ReadString(':')
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println("ReadString(':'):", str)

	// 使用 Seek 读取
	buf.Reset()
	buf.WriteString("try to SEEK a letter after colon :ABCDEFGHI")
	str1 := buf.Bytes()
	reader := bytes.NewReader(str1)

	if _, err := reader.Seek(30, io.SeekStart); err != nil {
		log.Println("reader Seek error:", err)
		return
	}
	// 读取 5 个字节
	readBuf := make([]byte, 5)
	n, err := reader.Read(readBuf)
	if err != nil {
		log.Println("reader Read error:", err)
		return
	}
	fmt.Println("reader Seek 后读取：", string(readBuf[:n]))

	// 读取 emoji
	buf.Reset()
	buf.WriteString("try to cut the emoji : (●'◡'●)")
	str1 = buf.Bytes()
	emoji := make([]byte, 36)
	reader = bytes.NewReader(str1)
	n, err = reader.Read(emoji)
	if err != nil {
		log.Println("reader Read emoji error:", err)
		return
	}
	fmt.Println("emoji 部分：", string(emoji[23:n]))


}

