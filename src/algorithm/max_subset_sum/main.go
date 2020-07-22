package main

import(
	// "time"
	"fmt"
	// "math/rand"
)

func main(){
	// //获取指定数量的整数序列
	// fmt.Println("请输入需要测试的数字序列长度：")
	// var len int
	// fmt.Scanf("%d",&len)
	// s := make([]int,len,len)
	// rand.Seed(time.Now().UnixNano())
	// // s := rand.Perm(len)
	// var char int
	// for n := 0;n < len;n++{
	// 	s[n] = rand.Intn(100)
	// 	char = rand.Intn(2)
	// 	if( char == 1 ){
	// 		s[n] = 0 - s[n]
	// 	}
	// }
	// fmt.Println(s)

	// 在未排序的数组中找到第 k 个最大的元素
	slice := []int{3,2,3,1,2,4,5,5,6}
	num := findKthLargest( slice,4)
	fmt.Println(num)
}

// 时间复杂度O(n^3)
func MaxSubsetSum1( s1 []int ){
	n1 := len(s1)
	fmt.Println(n1)
}

func findKthLargest(nums []int, k int) int {
    length := len(nums)
    start := 0
    end := length - 1
    index := partition(nums,start,end)
    for ;index != length - k;{
      //第length - k 在数组左边，继续划分数组左边
      if( length - k < index ){
      	end = index-1
        index = partition( nums,start,end )
      }else{
        //继续划分数组右边
        start = index+1
        index = partition( nums,start,end )
      }
    }
    return nums[index]
}

func partition( nums []int,start,end int )int{
		if( end == start  ){
			return start
		}
    middle_index := (end - start + 1)/2 + start
    //与最后一位交换顺序
    nums[middle_index],nums[end] = nums[end],nums[middle_index]
    last := end
    end--

    //作比较
    for ;start <= end;{
      if(nums[start] < nums[last] || nums[end] > nums[last] ){
      	if(nums[start] < nums[last]){
          start += 1
      	}
      	if( nums[end] > nums[last] ){
      		end -= 1
      	}
      }else{
	      if( nums[start] >= nums[last] &&  nums[end] <= nums[last] ){
	          nums[start],nums[end] = nums[end],nums[start]
	          start += 1
	          end -= 1
	      }
      }
    }
    nums[start],nums[last] = nums[last],nums[start]

    return start
}

// func findKthLargest(nums []int, k int) int {
//     // rand.Seed(time.Now().UnixNano())
//     return quickSelect(nums, 0, len(nums)-1, len(nums)-k)
// }

// func quickSelect(a []int, l, r, index int) int {
//     q := randomPartition(a, l, r)
//     fmt.Println(a)
//     fmt.Println("l",l)
//     fmt.Println("r",r)
//     fmt.Println("index",q)
//     if q == index {
//         return a[q]
//     } else if q < index {
//         return quickSelect(a, q + 1, r, index)
//     }
//     return quickSelect(a, l, q - 1, index)
// }

// func randomPartition(a []int, l, r int) int {
//     // i := rand.Int() % (r - l + 1) + l
//     i := (r - l + 1)/2 + l
//     a[i], a[r] = a[r], a[i]
//     return partition(a, l, r)
// }

// func partition(a []int, l, r int) int {
//     x := a[r]
//     i := l - 1
//     for j := l; j < r; j++ {
//         if a[j] <= x {
//             i++
//             a[i], a[j] = a[j], a[i]
//         }
//     }
//     a[i+1], a[r] = a[r], a[i+1]
//     return i + 1
// }