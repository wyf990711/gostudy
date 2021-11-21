package main

import (
	"fmt"
)

// 字符串
// go语言中字符串是用双引号包裹的
// 单引号包裹的是字符
func main() {

	s := "hello,你好"
	c1 := 'h'
	c2 := '王'
	c3 := "王"
	fmt.Println(s, c1, c2, c3)
	// 一个字节=8个bit （8个二进制位）
	// 一个字符‘A'=1个字节
	//一个utf8编码的汉字 ’王‘=一般占3个字节
	s1 := `
	第一行
第二行
第三行
	`
	fmt.Println(s1)
}
