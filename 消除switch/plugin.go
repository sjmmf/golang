package main

import "fmt"

// 两数相乘
func Multiply(x, y int) int {
	fmt.Println("请输入要进行计算的两个数:")
	fmt.Scanln(&x, &y)
	return x * y
}

// 两数相加
func Add(x, y int) int {
	fmt.Println("请输入要进行计算的两个数:")
	fmt.Scanln(&x, &y)
	return x + y
}

// 两数相减
func Subtract(x, y int) int {
	fmt.Println("请输入要进行计算的两个数:")
	fmt.Scanln(&x, &y)
	return x - y
}

// 两数相除
func Divide(x, y int) int {
	fmt.Println("请输入要进行计算的两个数:")
	fmt.Scanln(&x, &y)
	return x / y
}
