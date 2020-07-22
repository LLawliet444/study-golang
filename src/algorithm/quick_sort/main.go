package main

import(
	// "time"
	"fmt"
	// "math/rand"
)

func main(){
	// 快速排序
	slice := []int{2, 5, 3, 7, 4, 5, 8, 1, 4, 0}
	quickSort( slice)
	fmt.Println(slice)
}

func quickSort( nums []int ){
  length := len(nums)
  start := 0
  end := length - 1
  quickSortContent(nums, start, end)
}

func quickSortContent(nums []int, start, end int) {
  //边界条件
  if( end - start <= 0 ){
    return
  }
  index := partition( nums,start,end )
  //递归排左边
  quickSortContent( nums,start,index-1 )
  //递归排右边
  quickSortContent( nums,index+1,end )
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