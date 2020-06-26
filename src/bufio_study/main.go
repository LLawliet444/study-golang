package main

import(
	"fmt"
	"bufio"
	"io"
	"os"
)

func main(){
	// bufio通过缓冲来提高效率
	fileName := "/Users/xinxinzhang/code/go/test1.txt"
	file,err := os.Open(fileName)
	if err != nil{
		fmt.Println(err)
		return
	}
	defer file.Close()

	// // 创建Reader对象
	// b1 := bufio.NewReader(file)
	// // Read()
	// p := make([]byte,1024,1024)
	// n1,err := b1.Read(p)
	// fmt.Println(n1)
	// fmt.Println(string(p[:n1]))

	// // ReadLine()
	// data,flag,err := b1.ReadLine()
	// if err != nil{
	// 	fmt.Println(err)
	// 	return
	// }
	// fmt.Println(data)
	// fmt.Println(flag)
	// fmt.Println(string(data))

	// // ReadString()
	// for{
	// 	s1,err := b1.ReadString('\n')
	// 	if err == io.EOF{
	// 		fmt.Println("读取完毕")
	// 		return
	// 	}
	// 	fmt.Printf("%s",s1)
	// }

	// // ReadBytes()
	// data,err := b1.ReadBytes('\n')
	// if err != nil{
	// 	fmt.Println(err)
	// 	return
	// }
	// fmt.Println(string(data))

	// // Scanner
	// s2 := ""
	// fmt.Scanln(&s2)
	// fmt.Println(s2)

	// b2 := bufio.NewReader(os.Stdin)
	// s2,_ := b2.ReadString('\n')
	// fmt.Println(s2)

	file2,err := os.OpenFile(fileName,os.O_CREATE|os.O_WRONLY,os.ModePerm)
	if err != nil{
		fmt.Println(err)
		return
	}
	defer file2.Close()

	file2.Seek(0,io.SeekEnd)
	w1 := bufio.NewWriter(file2)
	// n,err := w1.WriteString("mother fuck")
	// fmt.Println(n)

	// // 刷新缓冲区（将缓冲区内容写到目标文件）
	// w1.Flush()

	// 缓冲器可用容量不足以放下p，且缓冲区有内容，用p把缓冲区填满，把缓冲区所有内容写入文件，并清空缓冲区
	for i := 1;i <= 1000;i++{
		w1.WriteString(fmt.Sprintf("%d:hello",i))
	}
	// 最后一次写入缓冲的内容没有满足写入文件并清空缓冲区的条件，需要手动刷新
	w1.Flush()

}