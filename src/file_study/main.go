package main

import(
	"fmt"
	"os"
	// "path/filepath"
	// "path"
)

func main(){
	// fileInfo,err := os.Stat("/Users/xinxinzhang/code/go/test1.txt")
	// if err != nil{
	// 	fmt.Println("err：",err)
	// 	return
	// }
	// fmt.Printf("%T\n",fileInfo)

	// // 获取文件信息
	// fmt.Println(fileInfo.Name())
	// fmt.Println(fileInfo.Size())
	// fmt.Println(fileInfo.IsDir())
	// fmt.Println(fileInfo.ModTime())
	// fmt.Println(fileInfo.Mode())

	// // 获取路径
	// fileName1 := "/Users/xinxinzhang/code/go/test1.txt"
	// fileName2 := "b.txt"

	// fmt.Println(filepath.IsAbs(fileName1))
	// fmt.Println(filepath.IsAbs(fileName2))
	// fmt.Println(filepath.Abs(fileName1))
	// fmt.Println(filepath.Abs(fileName2))

	// fmt.Println("获取父目录：",path.Join(fileName1,".."))

	// // 创建目录
	// err := os.Mkdir("/Users/xinxinzhang/code/go/fuck",os.ModePerm)
	// if err != nil{
	// 	fmt.Println("err:",err)
	// 	return
	// }
	// fmt.Println("创建fuck成功")

	// // 创建多级目录
	// err := os.MkdirAll("/Users/xinxinzhang/code/go/cc/dd",os.ModePerm)
	// if err != nil{
	// 	fmt.Println("err:",err)
	// 	return
	// }
	// fmt.Println("创建dd成功")

	// // 创建文件(权限为666，若文件存在，则截断)
	// file1,err := os.Create("/Users/xinxinzhang/code/go/fuck/a.txt")
	// if err != nil{
	// 	fmt.Println("err:",err)
	// 	return
	// }
	// fmt.Println(file1)

	// file2,err := os.Create(fileName2)
	// if err != nil{
	// 	fmt.Println("err:",err)
	// 	return
	// }
	// fmt.Println(file2)

	// // 打开文件
	// file3,err := os.Open(fileName1)
	// if err != nil{
	// 	fmt.Println("err:",err)
	// 	return
	// }
	// fmt.Println(file3)

	// file4,err := os.OpenFile(fileName1,os.O_RDONLY|os.O_WRONLY,os.ModePerm)
	// if err != nil{
	// 	fmt.Println("err:",err)
	// 	return
	// }
	// fmt.Println(file4)

	// // 关闭文件
	// file4.Close()

	// 删除文件或空文件夹
	// err := os.Remove("/Users/xinxinzhang/code/go/cc")
	// if err != nil{
	// 	fmt.Println("err:",err)
	// 	return
	// }
	// fmt.Println("删除文件夹成功")

	// 删除所有
	// err := os.RemoveAll("/Users/xinxinzhang/code/go/cc")	//文件夹不存在即成功
	// if err != nil{
	// 	fmt.Println("err:",err)
	// 	return
	// }
	// fmt.Println("删除文件夹成功")


}