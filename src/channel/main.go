package main

import (
	"fmt"
	"time"
)

func main(){
	ch1 := make(chan int)

	go func(){
		fmt.Println("子goroutine开始执行")
		time.Sleep(3*time.Second)
		data := <-ch1
		fmt.Println("data:",data)
	}()

	ch1 <- 10
	fmt.Println("main...over...")

	ch := make(chan int)
	ch <- 100
}