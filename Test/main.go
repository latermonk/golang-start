package main

import (
	"runtime"
	"time"
)

//func main() {
//
//}


func main() {
	runtime.GOMAXPROCS(1)
	for i := 0; i < 10; i++ {
		go println(i) //协程
		//println(i)
	}
	runtime.Gosched()
	time.Sleep(time.Second)
}