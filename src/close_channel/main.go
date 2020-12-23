package main

import (
	"fmt"
	"time"
)

func main(){
	ch1 := make(chan int)

	go sendData(ch1)

	// for{
	// 	time.Sleep(1*time.Second)
	// 	v,ok := <- ch1
	// 	if !ok{
	// 		fmt.Println("已经读取了所有的数据")
	// 		break
	// 	}
	// 	fmt.Println("读了：",v)
	// }

	for v := range ch1{
		time.Sleep(1*time.Second)
		fmt.Println("读取数据：",v)
	}
	fmt.Println("main...over...")
}

func sendData(ch1 chan int){
	for i := 0;i<10;i++{
		ch1 <- i
	}
	close(ch1)
}