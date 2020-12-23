package main

import (
	"fmt"
	"sync"
	"time"
)
var ticket = 10

// 创建🔐
var rwMutex *sync.RWMutex
var wg *sync.WaitGroup
func main(){
	rwMutex = new(sync.RWMutex)
	wg = new(sync.WaitGroup)

	wg.Add(3)

	// go readData(1)
	// go readData(2)

	go writeData(1)
	go readData(2)
	go readData(3)
	go writeData(4)

	wg.Wait()
	fmt.Println("main...over...")
}

func writeData( i int ){
	defer wg.Done()

	fmt.Println(i,"写")

	rwMutex.Lock()	//写操作上锁

	fmt.Println(i,"正在写数据。。。")

	time.Sleep(3*time.Second)

	rwMutex.Unlock()
	fmt.Println(i,"写结束。。。")
}

func readData(i int){
	defer wg.Done()

	fmt.Println(i,"读")

	rwMutex.RLock()	//读操作上锁

	fmt.Println(i,"正在读取数据。。。")

	time.Sleep(1*time.Second)

	rwMutex.RUnlock()
	fmt.Println(i,"读结束。。。")

}