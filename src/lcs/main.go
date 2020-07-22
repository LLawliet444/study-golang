package main

import (
	"fmt"
)
func main(){
	s := "abc"
	t := "ahbgdc"
	b := isSubsequence( s ,t)
	fmt.Println(b)
}

func isSubsequence(s string, t string) bool {
    //动态规划解题
    last_same_char_nums := make([]int,len(s),len(s))
    now_same_char_nums := make([]int,len(s),len(s))
    i := 0
    j := 0
    is_add := false
    for ;i<len(t);i++{
        j = 0
        is_add = false
        for ;j<len(s);j++{

            if( t[i] == s[j] ){
                if( j - 1 >= 0 ){
                    now_same_char_nums[j] = last_same_char_nums[j-1]+1
                    //等于s长度说明s为子序列
                    if( now_same_char_nums[j] == len(s) ){
                        return true
                    }

                }else{
                    now_same_char_nums[j] = 1
                }
                is_add = true
            }else{
                if( is_add ){
                    //本行有相似从左边取值
                    now_same_char_nums[j] = now_same_char_nums[j - 1]
                }else{
                    //本行无相似从上边取值
                    now_same_char_nums[j] = last_same_char_nums[j]
                }
            }
        }
        //结束后将now_same_char_nums值赋给last_same_char_nums
        // copy(last_same_char_nums,now_same_char_nums)
        for z := 0;z < len(now_same_char_nums);z++{
            last_same_char_nums[z] = now_same_char_nums[z]
        }
    }

    return false
}