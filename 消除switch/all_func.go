package main

import (
	"fmt"
	"plugin"
)

func main() {
	fmt.Println("======== Process Start ========")
	fmt.Println("加：Add         减：Subtract\n乘：Multiply     除：Divide")
	var func_name string
	var x int
	var y int
	fmt.Println("请选择一个功能:")
	fmt.Scanln(&func_name)
	// 查找插件文件
	pluginFile, err := plugin.Open("plugin.so")
	if err != nil {
		fmt.Println("An error occurred while opening the plug-in")
	} else {
		Func, err := pluginFile.Lookup(func_name)
		if err != nil {
			fmt.Println("没有实现此功能")
		} else {
			str := Func.(func(int, int) int)(x, y)
			fmt.Println("结果:", str)
		}
	}
}
