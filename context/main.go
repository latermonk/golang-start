package main

import (
	"fmt"
)


//func g(ctx context.Context){
func g(){
	fmt.Println("== Pig you are == ")

	//fmt.Println((ctx.Value("begin")))

}

func main() {
	fmt.Println("Hello")
	//ctx := context.WithValue(context.Background(), "begin", "From the Beginning")

	//go g(ctx)
	go g()
}
