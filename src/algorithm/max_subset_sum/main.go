package main

import(
	"time"
	"fmt"
	"math/rand"
)

func main(){
	//获取指定数量的整数序列
	fmt.Println("请输入需要测试的数字序列长度：")
	var len int
	fmt.Scanf("%d",&len)
	s := make([]int,len,len)
	rand.Seed(time.Now().UnixNano())
	// s := rand.Perm(len)
	var char int
	for n := 0;n < len;n++{
		s[n] = rand.Intn(100)
		char = rand.Intn(2)
		if( char == 1 ){
			s[n] = 0 - s[n]
		}
	}
	fmt.Println(s)
}

// 时间复杂度O(n^3)
func MaxSubsetSum1( s1 []int ){
	n1 := len(s1)
	fmt.Println(n1)
}