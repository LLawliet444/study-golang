package main

import(
	"fmt"
	"sort"
	"math"
	// "os"
	// "strconv"
)

/*
题目：输入1、2、3、4、5、6、7、8、9、0，将他们填入到完全二叉搜索树中，输出他们层序遍历的结果
思路：将数组从小到大排序，根节点为第左子树数量+1位数字，递归解决左子树和右子树
*/
func main(){
	arr := []int{1,2,3,4,5,6,7,8,9,0}
	result_arr := make([]int,10,10)
	sort.Ints(arr)
	solve(arr,0,result_arr)
	fmt.Println(result_arr)
}

func solve( arr []int,root_index int,result_arr []int ){
	if len(arr) == 0{
		return
	}
	// 获取左子树结点数
	n := len(arr);
	arr_root_index := getLeftLen(n)
	result_arr[root_index] = arr[arr_root_index]
	// 左树
	left_tree_arr := arr[:arr_root_index]
	// 层数
	// root_index/2+1
	solve(left_tree_arr,root_index*2+1,result_arr)
	//右树
	right_tree_arr := arr[arr_root_index+1:]

	solve(right_tree_arr,root_index*2+2,result_arr)
}

func getLeftLen(n int)int{
	// 高度
	float := float64(n+1)
	h := math.Log2( float)
	h = math.Ceil(h)
	x := n + 1 - int(math.Pow(2,h - 1.0))
	max_h := int(math.Pow(2,h - 1.0))/2
	if x > max_h {
		x = max_h
	}

	num := (int(math.Pow(2,h-1.0)) - 1)/2 + x
	return num
}