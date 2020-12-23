package main

import (
	"fmt"
)

func main(){
	var ch1 chan bool
	ch1 = make(chan bool)

	go func(){
		for i := 0;i < 10;i++{
			fmt.Println("子goroutine中，i：",i)
		}
		ch1 <- true
		fmt.Println("结束。。。")
	}()

	fmt.Println("主程序阻塞")
	data := <-ch1
	fmt.Println("main...data->",data)
	fmt.Println("main...over...")

}