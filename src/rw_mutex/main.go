package main

import (
	"fmt"
	"time"
	"math/rand"
	"sync"
)
var ticket = 10

// 创建🔐
var rwMutex sync.RWMutex
func main(){
	// a := 1
	// go func(){
	// 	a = 2
	// 	fmt.Println("goroutine中",a)
	// }()

	// a = 3
	// time.Sleep(1)
	// fmt.Println("main goroutine",a)

	wg.Add(4)
	// 4个售票口
	go saleTickets("售票口1")
	go saleTickets("售票口2")
	go saleTickets("售票口3")
	go saleTickets("售票口4")

	wg.Wait()

	// time.Sleep(10*time.Second)
}

func readData(i int){

	fmt.Println(i,"读")

	rwMutex.Rlock()


}