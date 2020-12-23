package main

import(
	"time"
	"fmt"
)

func main(){
	timer := time.NewTimer(5*time.Second)

	go func(){
		<- timer.C
		fmt.Println("定时器结束")
	}()

	// time.Sleep(3*time.Second)
	// flag := timer.Stop()
	// if flag{
	// 	fmt.Println("定时器中断")
	// }
}

