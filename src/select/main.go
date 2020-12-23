package main

import(
	"fmt"
	"time"
)
func main(){

	// ch1 := make(chan int)
	// ch2 := make(chan int)

	// go func(){
	// 	time.Sleep(3*time.Second)
	// 	ch2 <- 200
	// }()

	// go func(){
	// 	time.Sleep(3*time.Second)
	// 	ch1 <- 100
	// }()

	// select{
	// case num1 := <- ch1:
	// 	fmt.Println("通道1输出：",num1)
	// case num2,ok := <- ch2:
	// 	if ok{
	// 		fmt.Println("通道2输出：",num2)
	// 	}else{
	// 		fmt.Println("通道2关闭")
	// 	}
	// // default:
	// // 	fmt.Println("default")
	// }
	test1()
	fmt.Println("main...over...")
}

func test1(){
	ch1 := make(chan int)
	ch2 := make(chan int)

	go func(){
		ch1 <- 100
	}()

	select{
	case <- ch1:
		fmt.Println("case1执行")
	case <- ch2:
		fmt.Println("case2执行")
	case <- time.After(3*time.Second):
		fmt.Println("case3执行")
	default:
		fmt.Println("执行default")
	}
}
