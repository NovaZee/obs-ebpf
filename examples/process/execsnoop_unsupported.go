//go:build !linux

package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Printf("当前系统是 %s，不能加载 Linux eBPF 程序。\n", runtime.GOOS)
	fmt.Println("请在 Linux 环境中运行进程观测示例。")
}
