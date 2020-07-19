package main

import(
	"fmt"
	"sync"
)
var wg sync.WaitGroup
func main(){

	wg.Add(2)


	go fun1()
	go fun2()

	fmt.Println("main进入阻塞")
	wg.Wait()
	fmt.Println("main结束")
}

func fun1(){
	for i := 1;i<10;i++{
		fmt.Println("fun1()函数中打印",i)
	}
	wg.Done()	//counter -1
}
func fun2(){
	defer wg.Done()
	for j := 1;j<10;j++{
		fmt.Println("fun2()函数中打印",j)
	}

}

