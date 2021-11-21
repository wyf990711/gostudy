package main

import "fmt"

//基本数据类型-整型

func main() {
	var i1 = 10
	fmt.Printf("%d\n", i1) //十进制
	fmt.Printf("%b\n", i1) //八进制
	fmt.Printf("%o\n", i1) //八进制
	fmt.Printf("%x\n", i1) //十流进制
	//八进制
	i2 := 077 //63
	fmt.Printf("%d\n", i2)
	i3 := 0x123
	fmt.Printf("%d\n", i3)

	// %T查看变量类型
	fmt.Printf("%T\n", i3)
	// 声明int8类型
	i4 := int8(9) //明确指定声明int8类型，否者默认为int类型
	fmt.Printf("%T\n", i4)

}
