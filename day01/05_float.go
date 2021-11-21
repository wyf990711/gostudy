package main

import (
	"fmt"
)

func main() {
	f1 := 1.234
	fmt.Printf("%T", f1) //默认go语言中小数为flaot64
	f2 := float32(1.344)
	fmt.Printf("%T", f2)
}
