package main

import "fmt"

// 常量
// 定义常量之后不能修改，在程序运行期间不会改变
const Pi = 3.1415926

// 批量声明常量
const (
	statusOk = 200
	notFound = 402
)
const (
	// n2 n3 都是100
	n1 = 100  //100
	n2        //100
	n3        //100
	n4 = iota //3
)

// iota 在const关键字出现之后重置为0
const (
	a1 = iota //0
	a2 = 100  //100
	a3 = iota //2
	a4        //3
)
const (
	b1 = iota //0
	b2        //1
	_         //2
	b3        //3
)

// 定义数量级
const (
	_  = iota
	kB = 1 << (10 * iota) //二进制
	MB = 1 << (10 * iota)
	GB = 1 << (10 * iota)
	TB = 1 << (10 * iota)
	PB = 1 << (10 * iota)
)

func main() {
	fmt.Println("kb:", kB)
	fmt.Println("mb:", MB)
	fmt.Println("Gb:", GB)
	fmt.Println("Tb:", TB)
}
