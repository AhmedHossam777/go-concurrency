package main

import (
	"fmt"
	"time"
)

func someFunc(num string) {
	fmt.Println(num)
}

func main() {
	go someFunc("-90")
	go someFunc("090")
	go someFunc("190")

	time.Sleep(time.Second * 2)
	fmt.Println("hello, world")
}
