package main

import "fmt"

func loop() {
	for i := 0; i < 10; i++ {
		fmt.Printf("%d ", i)
	}
}

func main() {
	//go loop() // 启动一个goroutine
	loop()
}