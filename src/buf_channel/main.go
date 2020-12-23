package main

import(
	"fmt"
	"strconv"
	"time"
)

func main(){
	ch1 := make(chan string,4)

	go sendData(ch1)

	for{
		time.Sleep(1*time.Second)
		v,ok := <- ch1
		if !ok{
			fmt.Println("读取结束")
			break
		}
		fmt.Println("\t读取了：",v)
	}
	fmt.Println("main...over...")
}

func sendData( ch chan string ){
	for i := 0;i < 10;i++{
		ch <- "数据：" + strconv.Itoa(i)
		fmt.Printf("子goroutine写入了第%d个数据\n",i)
	}
	close(ch);
}

