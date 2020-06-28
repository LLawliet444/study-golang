package main

import(
	"fmt"
	"io/ioutil"
	// "os"
	// "strings"
	"log"
)

func main(){
	// 读取文件数据
	// fileName := "/Users/xinxinzhang/code/go/test1.txt"
	// data,err := ioutil.ReadFile(fileName)
	// if err != nil{
	// 	fmt.Println(err)
	// 	return
	// }
	// fmt.Println(string(data))

	// fileName := "/Users/xinxinzhang/code/go/test1.txt"
	// s1 := "hello world"
	// err := ioutil.WriteFile(fileName,[]byte(s1),os.ModePerm)
	// fmt.Println(err)

	// // ReadAll()
	// s2 := "mother fuck"
	// r1 := strings.NewReader(s2)
	// data,err := ioutil.ReadAll(r1)
	// if err != nil{
	// 	fmt.Println(err)
	// 	return
	// }
	// fmt.Println(string(data))

	// // ReadDir()
	// dirName := "/Users/xinxinzhang/code/go"
	// fileInfos,err := ioutil.ReadDir(dirName)
	// if err != nil{
	// 	fmt.Println(err)
	// 	return
	// }
	// for i:= 0;i<len(fileInfos);i++{
	// 	fmt.Printf("第%d个，名称：%s，是否是目录：%t\n",i,fileInfos[i].Name(),fileInfos[i].IsDir())
	// }
	// fmt.Println(fileInfos)

	// // 临时目录和临时文件
	// dir,err := ioutil.TempDir("/Users/xinxinzhang/code/go","Test")
	// if err != nil{
	// 	fmt.Println(err)
	// 	return
	// }
	// defer os.Remove(dir)
	// fmt.Println(dir)

	// file,err := ioutil.TempFile(dir,"Test")
	// if err != nil{
	// 	fmt.Println(err)
	// 	return
	// }
	// defer os.Remove(file.Name())
	// fmt.Println(file)

	// 递归遍历文件夹
	dirName := "/Users/xinxinzhang/code/go"
	listFiles(dirName,0)
}

func listFiles(dirName string,level int){
	s := "|--"
	for n := 0;n < level;n++{
		s = "|  "+s
	}
	fileInfos,err := ioutil.ReadDir(dirName)
	if err != nil{
		log.Fatal(err)
	}
	fileName := ""
	for i := 0;i < len(fileInfos);i++{
		fileName = dirName+"/"+fileInfos[i].Name()
		fmt.Printf("%s%s\n",s,fileName)
		if fileInfos[i].IsDir(){
			listFiles(fileName,level+1)
		}
	}
}