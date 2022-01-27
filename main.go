package main

import (
	"filterStr/src"
	"flag"
	"fmt"
)

func main()  {
	var t string
	var name string
	var length int
	flag.StringVar(&t, "t", "c", "类型,c:截取字符串, f:过滤字符串是否有效")
	flag.StringVar(&name, "n", "string", "字符串")
	flag.IntVar(&length, "l", 6, "长度")
	flag.Parse()

	if t == "c" {
		res := src.CutStr.CutString(name, length)
		fmt.Println(res)
	} else if t == "f" {
		err := src.FilterStr.FilStr(name)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(name, ": 字符串正常")
		}
	}
}
