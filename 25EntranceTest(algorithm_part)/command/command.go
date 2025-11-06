package main

import (
	"fmt"
	"strings"
)

const QUEUE_SIZE = 8

// 技能列表
var skillPatterns = map[string]string{
	"SDJ":   "波动拳",
	"DSDJ":  "升龙拳",
	"ASDL":  "防御反击",
	"SDJK":  "波动连击",
	"DSDJJ": "升龙连破",
	"WASJ":  "旋风拳",
	"ASDLK": "反击爆破",
	"SDJKL": "终极波动拳",
	"DJJJ":  "三连击",
}

func main() {
	var tempcmd string
	fmt.Scan(&tempcmd)

	//裁取后八位
	var cut string
	if len(tempcmd) > QUEUE_SIZE {
		cut = tempcmd[len(tempcmd)-9:]
	} else {
		cut = tempcmd
	}

	//转为大写
	cut = strings.ToUpper(cut)

	//检索特殊字符
	for _, v := range cut {
		if v == 'T' {
			cut = ""
			break
		}
	}

	if cut == "" {
		fmt.Println("已清空队列")
	} else {
		//去重
		for strings.Contains(cut, "AA") || strings.Contains(cut, "WW") || strings.Contains(cut, "SS") || strings.Contains(cut, "DD") {
			cut = strings.ReplaceAll(cut, "AA", "A")
			cut = strings.ReplaceAll(cut, "WW", "W")
			cut = strings.ReplaceAll(cut, "SS", "S")
			cut = strings.ReplaceAll(cut, "DD", "D")
		}

		skill, ok := skillPatterns[cut]

		if ok {
			fmt.Println(skill)
		}

		if !ok {
			fmt.Println("无匹配技能")
		}
	}
}
