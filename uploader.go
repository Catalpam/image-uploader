package main

import (
	"fmt"
	"os"
)

func main() {
	// 获取命令行参数
	var args = os.Args
	fmt.Println(args)
	if len(args) != 2 {
		return
	}
	return
}
