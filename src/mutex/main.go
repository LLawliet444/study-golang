package main

import (
	"fmt"
	"time"
	"math/rand"
	"sync"
)
var ticket = 10

// 创建🔐
var mutex sync.Mutex

var wg sync.WaitGroup
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

func saleTickets(name string){
	rand.Seed(time.Now().UnixNano())
	for{
		mutex.Lock()
		if ticket > 0{
			time.Sleep(time.Duration(rand.Intn(1000))*time.Millisecond)
			fmt.Println(name,"售出",ticket)
			ticket--
		}else{
			fmt.Println(name,"售罄")
			mutex.Unlock()
			break
		}
		mutex.Unlock()
	}
	wg.Done()
}