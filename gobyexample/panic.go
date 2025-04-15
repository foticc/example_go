package main

import "os"

func main() {
	// 输出一个错误消息和协程追踪信息，并以非零的状态退出程序。
	panic("Something went wrong")

	_, err := os.Create("/tmp/file.txt")
	if err != nil {
		panic(err)
	}
}
