package main

import(
	"fmt"
)

func main(){
	slice := []int{1,23,344,343,332,1,1,234,5,8}
	insertion_sort(slice)
	fmt.Println(slice)
}

// 从小到大排序
func insertion_sort(slice []int){
	i := 1
	for ;i < len(slice);i++{
		tmp := slice[i]
		j := i -1
		for ;j >= 0 && tmp < slice[j];j--{
			//小于则后挪一位
			slice[j+1] = slice[j]
		}
		slice[j+1] = tmp
	}
}