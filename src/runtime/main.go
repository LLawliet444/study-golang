package main

import(
	"fmt"
	"runtime"
	"time"
)

func main(){
	// fmt.Println("GOROOT->",runtime.GOROOT())
	// fmt.Println("os/platform->",runtime.GOOS)
	fmt.Println("逻辑CPU的数量->",runtime.NumCPU())
	// n := runtime.GOMAXPROCS(8)
	// fmt.Println(n)

	// Gosched
	// go func(){
	// 	for i:= 0;i<5;i++{
	// 		fmt.Println("goroutine...")
	// 	}
	// }()

	// for i:=0;i<4;i++{
	// 	runtime.Gosched()
	// 	fmt.Println("main...")
	// }

	go func(){
		fmt.Println("goroutine开始")
		fun()
		fmt.Println("goroutine结束")
	}()
	time.Sleep(3*time.Second)
}

func fun(){
	defer fmt.Println("defer...")
	runtime.Goexit()
	fmt.Println("fun函数")
}