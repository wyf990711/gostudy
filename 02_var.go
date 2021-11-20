package main

import "fmt"

// 函数外的每个语句都必须以const var func 等开头

func main() {
	// go语言中变量声明必须使用，不使用就编译不过去
	// var 变量名 类型
	var a int
	a = 3
	var b string
	b = "golang"
	fmt.Println(a, b)
	// 批量声明
	var (
		c int
		d int
	)
	c = 1
	d = 1
	fmt.Println(c, d)
	var (
		name string
		age  int
		isOk bool
	)
	name = "wyf"
	age = 23
	isOk = false
	fmt.Printf("name :%s,age:%d:,isok:%t\n", name, age, isOk)
	// 简短变量声明只能在函数内部使用
	s1 := "hello"
	fmt.Println(s1)
	// 当有些变量不想使用的时候 使用 _ 来忽略,相当于舍弃
	// s1 := 123
	// 同一个作用域中不能重复声明相同的变量
}
