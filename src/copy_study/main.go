package main

import(
	"fmt"
	"os"
	"io"
	"io/ioutil"
)

func main(){
	srcFile := "/Users/xinxinzhang/code/go/timg.jpeg"
	destFile := "copy3.jpeg"

	total,err := CopyFile3(srcFile,destFile)
	if err != nil{
		fmt.Println(err)
		return
	}
	fmt.Println(total,err)

	// 另外还有io.CopyN(),io.CopyBuffter()方法
}

// 不适用较大文件
func CopyFile3(srcFile,destFile string)(int,error){
	bs,err := ioutil.ReadFile(srcFile)
	if err != nil{
		return 0,err
	}
	err = ioutil.WriteFile(destFile,bs,0777)
	if err != nil{
		return 0,err
	}
	return len(bs),nil
}

func CopyFile2(srcFile,destFile string)(int64,error){
	file1,err := os.Open(srcFile)
	if err != nil{
		return 0,err
	}
	file2,err := os.OpenFile(destFile,os.O_WRONLY|os.O_CREATE,os.ModePerm)
	if err != nil{
		return 0,err
	}
	defer file1.Close()
	defer file2.Close()
	return io.Copy(file2,file1)
}

func CopyFile1(srcFile,destFile string)(int,error){
	file1,err := os.Open(srcFile)
	if err != nil{
		return 0,err
	}
	file2,err := os.OpenFile(destFile,os.O_WRONLY|os.O_CREATE,os.ModePerm)
	if err != nil{
		return 0,err
	}

	defer file1.Close()
	defer file2.Close()

	bs := make([]byte,1024,1024)
	n := -1
	total := 0
	for{
		n,err = file1.Read(bs)
		if err == io.EOF || n == 0{
			fmt.Println("拷贝完毕。")
			break
		}else if err != nil{
			fmt.Println("报错了")
			return total,err
		}
		total += n
		file2.Write(bs[:n])
	}

	return total,nil
}