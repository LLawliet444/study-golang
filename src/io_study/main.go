package main

import(
	"os"
	"fmt"
	// "io"
	"log"
)

func main(){
	// 读取文件数据
	// fileName := "/Users/xinxinzhang/code/go/test1.txt"
	// file,err := os.Open(fileName)
	// if err != nil{
	// 	fmt.Println("err:",err)
	// 	return
	// }
	// defer file.Close()

	// n := -1
	// bs := make([]byte,4,4)
	// for{
	// 	n,err = file.Read(bs)
	// 	if n == 0|| err == io.EOF{
	// 		fmt.Println("读取到文件末尾")
	// 		break
	// 	}
	// 	fmt.Println(string(bs[:n]))
	// }

	// 写数据
	fileName := "/Users/xinxinzhang/code/go/test2.txt"
	// file,err := os.Open(fileName)
	file,err := os.OpenFile(fileName,os.O_CREATE|os.O_WRONLY|os.O_APPEND,os.ModePerm)
	HandleErr(err)
	defer file.Close()

	// bs := []byte{65,66,67,68,69,70}
	bs := []byte{97,98,99,100}
	n,err := file.Write(bs[2:])
	HandleErr(err)
	fmt.Println(n)
	file.WriteString("\n")

	// 写字符串
	n,err = file.WriteString("hello world\n")
	HandleErr(err)
	fmt.Println(n)

	n,err = file.Write([]byte("today\n"))
	HandleErr(err)
	fmt.Println(n)

}

func HandleErr(err error){
	if err != nil{
		log.Fatal(err)
	}
}