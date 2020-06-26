package main

import(
	"io"
	"os"
	"fmt"
	"strings"
	"log"
	"strconv"
)

func main(){
	// Seek使用
	// fileName := "/Users/xinxinzhang/code/go/test1.txt"
	// file,err := os.OpenFile(fileName,os.O_RDWR,os.ModePerm)
	// if err != nil{
	// 	fmt.Println(err)
	// 	return
	// }
	// defer file.Close()
	// bs := []byte{0}
	// file.Read(bs)
	// fmt.Println(string(bs))

	// file.Seek(3,io.SeekStart)
	// file.Read(bs)
	// fmt.Println(string(bs))

	// file.Seek(2,io.SeekCurrent)
	// file.Read(bs)
	// fmt.Println(string(bs))

	// file.Seek(0,io.SeekEnd)
	// file.WriteString("ABC")

	// 断点续传
	srcFile := "/Users/xinxinzhang/code/timg.jpeg"
	destFile := srcFile[strings.LastIndex(srcFile,"/")+1:]
	tempFile := destFile+"temp.txt"

	file1,err := os.Open(srcFile)
	HandleErr(err)
	file2,err := os.OpenFile(destFile,os.O_CREATE|os.O_WRONLY,os.ModePerm)
	HandleErr(err)
	file3,err := os.OpenFile(tempFile,os.O_CREATE|os.O_RDWR,os.ModePerm)
	HandleErr(err)

	defer file1.Close()
	defer file2.Close()

	// 先读取临时文件中的数据，再seek
	file3.Seek(0,io.SeekStart)
	bs := make([]byte,100,100)
	n1,err := file3.Read(bs)
	HandleErr(err)
	countStr := string(bs[:n1])
	count,err := strconv.ParseInt(countStr,10,64)		//??why
	// count,err = strconv.ParseInt("abcddd",10,64)
	// HandleErr(err)
	// fmt.Println(count)

	file1.Seek(count,io.SeekStart)
	file2.Seek(count,io.SeekStart)
	data := make([]byte,1024,1024)
	n2 := -1
	n3 := -1
	total := int(count)
	for{
		n2,err = file1.Read(data)
		if err == io.EOF || n2 == 0{
			fmt.Println("文件复制结束",total)
			file3.Close()
			os.Remove(tempFile)
			break
		}

		n3,err = file2.Write(data[:n2])
		total += n3

		file3.Seek(0,io.SeekStart)
		file3.WriteString(strconv.Itoa(total))
		fmt.Println("total:",total)

		// if total > 8000{
		// 	panic("断电了")
		// }
	}
}

func HandleErr(err error){
	if err != io.EOF && err != nil{
		log.Fatal(err)
	}
}