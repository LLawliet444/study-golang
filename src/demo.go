package main

import(
	"fmt"
	// "sort"
	// "strings"
	// "strconv"
)

func main(){
	// // 51 数组（值类型）
	// arr := [4]int{1,2,3,4}
	// fmt.Println(arr)

	// // 切片（引用类型）
	// // 切片存储的是底层数组的地址
	// // 切片扩容成倍增长
	// // 切片每次扩容都会改变数组的地址
	// s1 := []int{1,2,3,4}
	// fmt.Println(s1)
	// fmt.Printf("%T,%T\n",arr,s1)

	// // make方法创建切片
	// s2 := make([]int,3,8)
	// fmt.Printf("长度：%d，容量：%d\n",len(s2),cap(s2))
	// fmt.Println(s2)
	// // fmt.Println(s2[3])		//panic: runtime error: index out of range [3] with length 3
	// s2 = append(s2,1,2)		//append从长度末尾添加，容量不够自动扩容
	// fmt.Println(s2)
	// s2 = append(s2,s1...)		//...意为添加s1中的元素
	// fmt.Println(s2)


	// // 51 已有数组直接创建切片
	// // 从数组创建切片为引用，修改任意一个元素别的变量中的这个元素都会变化
	// // 切片扩容会指向新的数组
	// a := [10]int{1,2,3,4,5,6,7,8,9,10}
	// fmt.Println(a)
	// s1 := a[:5]
	// s2 := a[3:8]
	// s3 := a[5:]
	// s4 := a[:]
	// fmt.Println("s1:",s1)
	// fmt.Println("s2:",s2)
	// fmt.Println("s3:",s3)
	// fmt.Println("s4:",s4)

	// // 52 值类型为值传递，即为深拷贝。引用类型为引用传递，拷贝地址，即为浅拷贝，深拷贝可以手动拷贝或者使用copy()
	// // copy()不能自动扩容
	// s1 := []int{1,2,3,4}
	// s2 := make([]int,2,4)
	// copy(s2,s1[2:])
	// fmt.Println(s1)
	// fmt.Println(s2)

	// // 54 map
	// // map无序
	// // map长度不固定，为引用类型
	// // map的key必须为可以比较的类型
	// // 引用类型初始值为nil
	// // 创建map
	// var map1 map[int]string
	// var map2 = make(map[int]string)
	// var map3 = map[string]int{"go":90,"php":40,"java":98,"js":87}
	// fmt.Println(map1)
	// fmt.Println(map2)
	// fmt.Println(map3)
	// // nil map
	// fmt.Println(map1 == nil)
	// fmt.Println(map2 == nil)
	// fmt.Println(map3 == nil)
	// if( map1 == nil ){
	// 	map1 = make(map[int]string)
	// }

	// // 存储键值对到map
	// map1[1] = "林克"
	// map1[2] = "吉良吉影"
	// map1[3] = "花择言"
	// map1[4] = "顾惜朝"
	// // key存在获取数值，key不存在获取的为0值
	// fmt.Println(map1) 

	// v1,ok := map1[40]
	// if ok{
	// 	fmt.Println("数值为：",v1)
	// }else{
	// 	fmt.Println("key不存在")
	// }

	// // 修改数据
	// map1[1] = "fuck mother"
	// fmt.Println(map1)

	// // 删除数据(key存在删除，key不存在无影响)
	// delete(map1,1)
	// fmt.Println(map1)

	// // 长度
	// fmt.Println(len(map1))

	// // 遍历 for range
	// for k,v := range map1{
	// 	fmt.Println(k,v)
	// }

	// // 顺序遍历
	// keys := make([]int,0,len(map1))
	// for k,_ := range map1{
	// 	keys = append(keys,k)
	// }
	// fmt.Println(keys)
	// sort.Ints(keys)
	// fmt.Println(keys)

	// for _,key := range keys{
	// 	fmt.Println(key,map1[key])
	// }

	// s := make([]map[int]string,0,2)
	// s = append(s,map1)
	// s = append(s,map2)
	// fmt.Println(s)
	// // 引用类型
	// s[1][2] = "塞尔达"
	// fmt.Println(s)
	// fmt.Println(map2)

	// // 59 string
	// s1 := "hello 中国"
	// s2 := "hello world"
	// for _,v := range s2 {
	// 	fmt.Printf("%c",v)
	// }
	// fmt.Println("\n")
	// for _,v := range s1 {
	// 	fmt.Printf("%c",v)
	// }

	// // 字节切片转字符串
	// slice1 := []byte{65,66,67,68,69}
	// s3 := string(slice1)
	// fmt.Println(s3)

	// // 字符串转切片
	// s4 := "abcde"
	// slice2 := []byte(s4)
	// fmt.Println(slice2) 

	// // 字符串不能修改

	// // strings包
	// fmt.Println(strings.Contains(s1,"e"))
	// fmt.Println(strings.ContainsAny(s1,"abcd"))
	// fmt.Println(strings.Count(s1,"l"))
	// fmt.Println(strings.HasPrefix(s1,"h"))
	// fmt.Println(strings.HasSuffix(s1,"中国"))
	// fmt.Println(strings.Index(s1,"中国"))
	// fmt.Println(strings.IndexAny(s1,"中国l"))
	// fmt.Println(strings.LastIndex(s1,"l"))
	// fmt.Println(strings.LastIndexAny(s1,"中国l"))
	// // 字符串拼接
	// ss1 := []string{"hello","world"}
	// s5 := strings.Join(ss1,",")
	// fmt.Println(s5)
	// // 字符串切割
	// ss2 := strings.Split(s5,",")
	// fmt.Println(ss2)
	// // 截取字符串
	// s6 := s1[:5]
	// fmt.Println(s6)

	// // 字符串转换
	// s7 := "true"
	// b1,err := strconv.ParseBool(s7)
	// if err != nil{
	// 	fmt.Println(err)
	// 	return
	// }
	// fmt.Printf("%T\n",b1)
	// ss3 := strconv.FormatBool(b1)
	// fmt.Printf("%T\n",ss3)
	// i3,err := strconv.Atoi("-20")
	// fmt.Printf("%T,%d",i3,i3)

	// 62 函数（值类型值传递，引用类型引用传递）
	sum1 := getSum(10);
	fmt.Println(sum1)
	getAdd(1,2)
	getSum1(1,2,3,4)
	s1 := []int{1,2,3,4,5}
	getSum1(s1...)

	res1,res2 := rectangle(3,5)
	fmt.Printf("周长：%f,面积：%f",res1,res2)
}

func getSum(n int)(sum int){
	for i := 1;i <= n; i++{
		sum += i
	}
	// fmt.Println(sum)
	return sum
}

func getAdd(a ,b int){
	sum := a+b
	fmt.Println(sum)
}

// 可变参数(传入的是切片，可变参数要放最后且只能有一个)
func getSum1(nums ... int){
	sum := 0
	for i := 0;i < len(nums);i++{
		sum  += nums[i]
	}
	fmt.Println(sum)
}

// 求矩形的周长和面积
func rectangle(len,wid float64)(float64,float64){
	perimeter := (len + wid)*2
	area := len * wid
	return perimeter,area
}