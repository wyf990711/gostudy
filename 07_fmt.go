package main

import "fmt"

func main() {
	i := 100
	fmt.Printf("%T\n", i) //%T:类型
	fmt.Printf("%v\n", i) //%v:值
	fmt.Printf("%b\n", i) //%b:二进制
	fmt.Printf("%d\n", i) //%d:十进制
	fmt.Printf("%o\n", i) //%o:八进制
	fmt.Printf("%x\n", i) //%x:十六进制
	s := "www"
	fmt.Printf("%s\n", s) // %s：字符串
}
