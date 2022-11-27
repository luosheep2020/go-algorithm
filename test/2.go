package main

import "fmt"

func main() {
	ch1 := make(chan string, 1)
	ch2 := make(chan string, 1)
	ch3 := make(chan string, 1)
	ch1 <- "Go语言微服务架构核心22讲"
	// 往通道 2 发送数据 从0到Go语言微服务架构师
	ch2 <- "从0到Go语言微服务架构师"
	// 往通道 3 发送数据 Go语言极简一本通
	ch3 <- "Go语言极简一本通"

	select {
	case message1 := <-ch1:
		fmt.Println("ch1 received:", message1)
	case message2 := <-ch2:
		fmt.Println("ch2 received:", message2)
	case message3 := <-ch3:
		fmt.Println("ch3 received:", message3)
	default:
		fmt.Println("No data received.")
	}
}
