package main

import(
	"fmt"
)

func main(){
	array := []int{1,3,5,6,2,4,5,3,4,4}
	bubble_sort(array)
	fmt.Println(array)
}

func bubble_sort( slice []int){
	num := 0
	i := 0
	for ;i + 1 < len(slice);i ++{
		j := 0
		flag := 0
		for ;j + 1 < len(slice) - i;j++ {
			if( slice[j] > slice[j+1] ){
				flag = 1
				num = slice[j+1]
				slice[j+1] = slice[j]
				slice[j] = num
			}
		}
		if flag == 0{
			break
		}
	}
}