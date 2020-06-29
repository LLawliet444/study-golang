package main

import(
	"fmt"
	"time"
)

func main(){
	go printNum()

	for i := 1;i<= 100;i++{
		fmt.Println("主goroutine中的字母A:",i)
	}
	time.Sleep(1*time.Second)
	fmt.Println("over")
}

func printNum(){
	for i := 1;i<= 100;i++{
		fmt.Println("子goroutine中的数字",i)
	}
}