package main

import(
	"fmt"
	"math"
	// "sort"
	// "strings"
	"strconv"
	// "os"
	// "log"
	// "errors"
	// "net"
	// "path/filepath"
	"l_package/utils/timeutils"
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

	// // 62 函数（值类型值传递，引用类型引用传递）
	// sum1 := getSum(10);
	// fmt.Println(sum1)
	// getAdd(1,2)
	// getSum1(1,2,3,4)
	// s1 := []int{1,2,3,4,5}
	// getSum1(s1...)

	// res1,res2 := rectangle(3,5)
	// fmt.Printf("周长：%f,面积：%f",res1,res2)

	// 递归函数
	// fmt.Println(getSum2(5))
	// fmt.Println(getFibonacci(8))
	// // 1 2 3 4 5 6 7 8
  // // 	1 1 2 3 5 8 13 21

  // // defer 延迟函数，等主函数执行完再执行，使用栈，先入后出
  // // defer函数调用时，已传参不执行，return在defer函数执行完之后执行
  // // 用法	临时文件的删除
  // 		// 异常处理 panic()和recover() recover()语法上要求在defer中执行
  // defer fun1("hello")
  // defer fun1("world")

  // fmt.Printf("%T\n",rectangle)

  // // 定义函数类型的变量
  // var c func(int)int
  // c = getSum
  // fmt.Println(c)
  // a := c(10)
  // fmt.Println(a)

  // // 匿名函数(通过小括号直接调用，通常只能使用一次)
  // // go支持函数式编程
  // // 	1.将匿名函数作为另一个函数的参数，即回调函数
  // // 	2.将匿名函数作为另一个函数的返回值，可以形成闭包结构
  // fun2 := func (){
  // 	fmt.Println("匿名函数")
  // }
  // fun2()
  // fun2()
  // // 定义带参数的匿名函数
  // func (a,b int){
  // 	fmt.Println(a,b)
  // }(1,2)
  // // 定义带返回值的匿名函数
  // var res = func (a,b int)(int){
  // 	return a + b
  // }(3,4)
  // fmt.Println(res)

 	// // 回调函数
 	// // fmt.Printf("%T\n%T\n",add,oper)

 	// res1 := add(1,2)
 	// fmt.Println(res1)
 	// res2 := oper(1,2,sub)
 	// fmt.Println(res2)
 	// fun1 := func(a,b int)int{
 	// 	return a*b
 	// }
 	// fmt.Println(oper(10,20,fun1))

 	// // 闭包
 	// res1 := increment()
 	// fmt.Printf("%T\n",res1)
 	// v1 := res1()
 	// fmt.Println(v1)
 	// v2 := res1()
 	// fmt.Println(v2)

 	// res2 := increment()
 	// v3 := res2()
 	// fmt.Println(v3)

 	// v4 := res1()
 	// fmt.Println(v4)

 	// // 69 指针
 	// a := 10
 	// fmt.Printf("%p\n",&a)
 	// var p1 *int
 	// fmt.Println(p1)
 	// p1 = &a
 	// fmt.Println(p1)
 	// fmt.Println(*p1)
 	// a = 100
 	// fmt.Printf("%p\n",&a)
 	// *p1 = 200
 	// fmt.Println(a)

 	// // 指针的指针
 	// var p2 **int
 	// fmt.Println(p2)
 	// p2 = &p1
 	// fmt.Printf("%T,%T,%T\n",a,p1,p2)
 	// **p2 = 300
 	// fmt.Println(a)

 	// // 数组指针（首先是指针）
 	// arr1 := [4]int{1,2,3,4}
 	// fmt.Println(arr1)
 	// var p1 *[4]int
 	// p1 = &arr1
 	// fmt.Println(p1)
 	// fmt.Println(*p1)
 	// (*p1)[0] = 100
 	// p1[1] = 200
 	// fmt.Println(p1)
 	// // 指针数组（首先是数组）
 	// a := 1
 	// b := 2
 	// c := 3
 	// d := 4
 	// arr2 := [4]int{a,b,c,d}
 	// arr3 := [4]*int{&a,&b,&c,&d}
 	// fmt.Println(arr2)
 	// fmt.Println(arr3)
 	// *arr3[0] = 100
 	// fmt.Println(arr2)
 	// fmt.Println(a)
 	// // 函数指针
 	// a := add
 	// fmt.Printf("%p\n",&a)
 	// // 指针函数
 	// arr2 := fun3()
 	// fmt.Printf("arr2:%T,%p,%v\n",arr2,arr2,arr2)
 	// // 指针作为参数（引用传递）
 	// a := 100
 	// fun2(&a)
 	// fmt.Println(a)

 	// // 结构体
 	// var p1 Person
 	// fmt.Println(p1)
 	// p1.name = "吉良吉影"
 	// p1.age = 2
 	// p1.sex = "女"
 	// p1.address = "猫窝"
 	// fmt.Printf("姓名：%s,年龄：%d,性别：%s,地址：%s\n",p1.name,p1.age,p1.sex,p1.address)
 	// p2 := Person{}
 	// p2.name = "林克"
 	// fmt.Printf("姓名：%s,年龄：%d,性别：%s,地址：%s\n",p2.name,p2.age,p2.sex,p2.address)
 	// p3 := Person{name:"花择言",age:1,sex:"男",address:"花家地"}
 	// fmt.Printf("姓名：%s,年龄：%d,性别：%s,地址：%s\n",p3.name,p3.age,p3.sex,p3.address)
 	// p4 := Person{"李小新",18,"女","望京"}
 	// fmt.Printf("姓名：%s,年龄：%d,性别：%s,地址：%s\n",p4.name,p4.age,p4.sex,p4.address)
 	// // 结构体值传递
 	// p5 := p1
 	// fmt.Printf("%p,%p\n",&p1,&p5)
 	// // 定义结构体指针（引用传递）
 	// var pp1 *Person
 	// pp1 = &p1
 	// pp1.name = "李四"
 	// fmt.Println(p1)
 	// // new返回指针，且不为nil
 	// pp2 := new(Person)
 	// fmt.Printf("%T\n",pp2)
 	// pp2.name = "Jerry"
 	// pp2.age = 20
 	// pp2.sex = "男"
 	// pp2.address = "美国"
 	// fmt.Println(pp2)

 	// // 匿名结构体
 	// p1 := struct{
 	// 	name string
 	// 	age int
 	// }{
 	// 	name:"李四",age:19,
 	// }
 	// fmt.Println(p1)

 	// // 匿名字段
 	// w1 := Worker{"李晓华",10}
 	// fmt.Println(w1.string)

 	// // 嵌套结构体
 	// b1 := Book{"西游记",45.9}
 	// s1 := Student{"王二狗",18,b1}
 	// fmt.Println(s1)
 	// fmt.Println(s1.book.bookName)
 	// s2 := Student{"李小花",18,Book{"娱乐至死",28}}
 	// fmt.Println(s2)
 	// s1.book.bookName = "红楼梦"
 	// fmt.Println(b1)

 	// b2 := Book{"呼啸山庄",30}
 	// s3 := Student2{"小心",29,&b2}
 	// fmt.Println(s3)
 	// s3.book.bookName = "雾都孤儿"
 	// fmt.Println(b2)

 	// // 面向对象（提升字段）
 	// var s1 Student3
 	// s1.name = "吉良吉影"
 	// s1.age = 2
 	// s1.school = "清华大学"
 	// fmt.Println(s1)

 	// // 方法
 	// w1 := Worker{name:"王二狗",age:30,sex:"男"}
 	// w1.work()
 	// w2 := &Worker{name:"Ruby",age:32,sex:"女"}
 	// w2.work()
 	// w2.rest()
 	// w1.rest()
 	// w1.printInfo()
 	// c1 := Cat{color:"蓝白",age:1}
 	// c1.printInfo()
 	// // 方法的继承和改写
 	// p1 := Person{name:"王二狗",age:30}
 	// p1.eat()
 	// s1 := Student3{Person:Person{name:"吉良吉影",age:10,sex:"女",address:"望京"},school:"清华大学"}
 	// s1.eat()
 	// s1.study()

 	// // 接口(方法必须全部实现，非侵入式)
 	// m1 := Mouse{name:"罗技小红"}
 	// f1 := FlashDisk{name:"闪迪64G"}
 	// testInterFace(m1)
 	// testInterFace(f1)
 	// var usb USB
 	// usb = m1
 	// usb.start()
 	// f1.delete()
 	// var arr[2]USB
 	// arr[0] = m1
 	// arr[1] = f1
 	// fmt.Println(arr)

 	// // 空接口(可以存储任意类型的数据)，可以作为容器
 	// var a1 A = 100
 	// var a2 A = "mother fuck"
 	// test1(a1)
 	// test1(a2)
 	// map1 := make(map[string]interface{})
 	// map1["name"] = "李晓华"
 	// map1["age"] = 30
 	// map1["friend"] = Person{name:"梨花",age:10}
 	// fmt.Println(map1)

 	// slice1 := make([]interface{},0,10)
 	// slice1 = append(slice1,a1,a2,"吉良吉影",map1)
 	// fmt.Println(slice1 )

 	// var cat Cat = Cat{color:"狸花"}
 	// cat.test1()
 	// cat.test2()
 	// cat.test3()

 	// // 接口嵌套
 	// var a1 A1 = cat
 	// a1.test1()
 	// var b1 B = cat
 	// b1.test2()
 	// var c1 C = cat
 	// c1.test1()
 	// c1.test2()
 	// c1.test3()

 	// var a2 A1 = c1
 	// a2.test1()

 	// // 接口断言
 	// var t1 Triangle = Triangle{3,4,5}
 	// fmt.Println(t1.peri())
 	// fmt.Println(t1.area())

 	// var c1 Circle = Circle{4}
 	// fmt.Println(c1.peri())
 	// fmt.Println(c1.area())

 	// var s1 Shape
 	// s1 = t1
 	// fmt.Println(s1.peri())
 	// fmt.Println(s1.area())

 	// var s2 Shape = c1
 	// fmt.Println(s2.peri())
 	// fmt.Println(s2.area())

 	// testShape(t1)
 	// testShape(s1)
 	// testShape(s2)

 	// getType(t1)
 	// getType(c1)
 	// getType(s1)

 	// var t2 *Triangle = &Triangle{3,4,3}
 	// getType(t2)

 	// getType2(t1)
 	// getType2(c1)

 	// // Type关键词
 	// var i1 myint = 100
 	// var i2 = 200
 	// var name mystr = "王二狗"
 	// var s1 = "mother fuck"
 	// fmt.Println(i1,i2)
 	// fmt.Println(name,s1)

 	// fmt.Printf("%T,%T,%T,%T\n",i1,i2,s1,name)

 	// res1 := fun10()
 	// fmt.Println(res1(10,20))
 	// var i3 myint2 = 300
 	// i2 = i3
 	// fmt.Println(i2)

 	// // ERROR
 	// f,err := os.Open("test.txt")
 	// if err != nil{
 	// 	// 错误类型
 	// 	fmt.Println(err)
 	// 	if ins,ok := err.(*os.PathError);ok{
 	// 		fmt.Println("1.Op:",ins.Op)
 	// 		fmt.Println("2.Path:",ins.Path)
 	// 		fmt.Println("3.Err:",ins.Err)
 	// 	}
 	// 	return
 	// }
 	// fmt.Println(f.Name(),"打开文件成功")

 	// addr,err := net.LookupHost("www.jintianddddd.com")
 	// if err != nil{
 	// 	fmt.Println(err)
 	// 	if ins,ok := err.(*net.DNSError);ok{
 	// 		if ins.Timeout(){
 	// 			fmt.Println("操作超时")
 	// 		}else if ins.Temporary(){
 	// 			fmt.Println("临时性错误")
 	// 		}else{
 	// 			fmt.Println("通常错误")
 	// 		}
 	// 	}
 	// 	return
 	// }
 	// fmt.Println(addr)

 	// files,err := filepath.Glob("[")
 	// if err != nil && err == filepath.ErrBadPattern{
 	// 	fmt.Println(err)
 	// 	return
 	// }
 	// fmt.Println("files:",files)

 	// // 创建错误
 	// err1 := errors.New("创建错误")
 	// fmt.Println(err1)
 	// fmt.Printf("%T\n",err1)
 	// err2 := fmt.Errorf("错误的信息码：%d",100)
 	// fmt.Println(err2)
 	// fmt.Printf("%T\n",err2)

 	// err3 := checkAge(-10)
 	// if err3 != nil{
 	// 	fmt.Println(err3)
 	// 	return
 	// }
 	// fmt.Println("没有错误")

 	// // 自定义错误
 	// var c Circle = Circle{radius:-10}
 	// a1,err := c.area()
 	// if err != nil{
 	// 	fmt.Println(err)
 	// 	if err,ok := err.(*areaError);ok{
 	// 		fmt.Printf("半径是：%2f\n",err.radius)
 	// 	}
 	// 	return
 	// }
 	// fmt.Println(a1)

 	// length,width := 3.4,-9.8
 	// area,err := rectArea(length,width)
 	// if err != nil{
 	// 	fmt.Println(err)
 	// 	if err,ok := err.(*rectAreaError);ok{
 	// 		if err.lengthNegative(){
 	// 			fmt.Printf("error:长度%2f小于0",err.length)
 	// 		}
 	// 		if err.widthNegative(){
 	// 			fmt.Printf("error:宽度%2f小于0",err.width)
 	// 		}
 	// 	}
 	// 	return
 	// }
 	// fmt.Println(area)

 // 	// panic()和recover() 适用：空指针引用、下标越界、除数为0、不应该出现的分支、输入不应该引起函数错误
 // 	defer func(){
	// 	if msg := recover();msg != nil{
	// 		fmt.Println(msg,"程序回复")
	// 	}
	// }()

 // 	funA()
 // 	defer myprint("defer main:3")
 // 	funB()
 // 	defer myprint("defer main:4")
 // 	fmt.Println("over")

 timeutils.PrintTime()

}
func myprint(s string){
	fmt.Println(s)
}
func funA(){
	fmt.Println("我是函数A")
}
func funB(){
	fmt.Println("我是函数B")
	defer myprint("defer funB():1")
	for i:= 1;i<=10;i++{
		fmt.Println(i)
		if i == 5{
			panic("funB函数panic")
		}
	}
	defer myprint("defer funB():2")
}
type rectAreaError struct{
	msg string
	length float64
	width float64
}
func (e *rectAreaError)Error()string{
	return e.msg
}
func (e *rectAreaError)lengthNegative()bool{
	return e.length < 0
}
func (e *rectAreaError)widthNegative()bool{
	return e.width < 0
}
func rectArea(length,width float64)(float64,error){
	msg := ""
	if length < 0{
		msg = "长度小于0"
	}
	if width < 0{
		if msg == ""{
			msg = "宽度小于0"
		}else{
			msg += ",宽度也小于0"
		}
	}
	if msg != ""{
		return 0,&rectAreaError{msg,length,width}
	}
	return length*width,nil
}
func checkAge(age int)error{
	if age < 0{
		// return errors.New("年龄不合法")
		err := fmt.Errorf("给定的年龄%d不合法",age)
		return err
	}
	fmt.Println("年龄为：",age)
	return nil
}
type areaError struct{
	msg string
	radius float64
}
func (e *areaError)Error()string{
	return fmt.Sprintf("error:半径,%2f,%s",e.radius,e.msg)
}
// 定义新类型
type myint int
type mystr string
// 定义函数类型
type myfun func(int,int)(string)
func fun10()myfun{
	fun := func(a,b int)string{
		s := strconv.Itoa(a)+strconv.Itoa(b)
		return s
	}
	return fun
}
// 类型别名
type myint2 = int
func getType(s Shape){
	if ins,ok := s.(Triangle);ok{
		fmt.Println("是三角形，三边为：",ins.a,ins.b,ins.c)
	}else if ins,ok := s.(Circle);ok{
		fmt.Println("是圆形，半径是：",ins.radius)
	}else{
		fmt.Println("我也不知道了")
	}
}
func getType2(s Shape){
	switch ins := s.(type){
	case Triangle:
		fmt.Println("三角形",ins.a,ins.b,ins.c)
	case Circle:
		fmt.Println("圆形",ins.radius)
	}
}
func testShape(s Shape){
	fmt.Printf("周长：%2f,面积：%2f\n",s.peri())
}
type Shape interface{
	peri() float64 //周长
	area() (float64,error) //面积
}
type Triangle struct{
	a,b,c float64
}
func (t Triangle)peri()float64{
	return t.a+t.b+t.c
}
func (t Triangle)area()(float64,error){
	if t.a < 0 {
		return 0,&areaError{"长度是非法的",t.a}
	}

	p := t.peri()/2
	s := math.Sqrt( p*(p-t.a)*(p-t.b)*(p-t.c) )
	return s,nil
}
type Circle struct{
	radius float64
}
func (c Circle)peri()float64{
	return c.radius*2*math.Pi
}
func (c Circle)area()(float64,error){
	if c.radius < 0 {
		return 0,&areaError{"半径是非法的",c.radius}
	}
	return math.Pow(c.radius,2)*math.Pi,nil
}
type A1 interface{
	test1()
}
type B interface{
	test2()
}
type C interface{
	A1
	B
	test3()
}
func(cat Cat)test1(){
	fmt.Println("test1")
}
func(cat Cat)test2(){
	fmt.Println("test2")
}
func(cat Cat)test3(){
	fmt.Println("test3")
}
func test1(a A){
	fmt.Println(a)
}
type A interface{

}
type USB interface{
	start() //USB设备开始工作
	end()	//USB设备结束工作
}
func testInterFace(usb USB){
	usb.start()
	usb.end()
}
type Mouse struct{
	name string
}
type FlashDisk struct{
	name string
}
func(m Mouse)start(){
	fmt.Println(m.name,"鼠标准备就绪，可以开始工作")
}
func(m Mouse)end(){
	fmt.Println(m.name,"鼠标结束工作，可以安全退出")
}
func(f FlashDisk)start(){
	fmt.Println(f.name,"准备开始工作，可以存储")
}
func(f FlashDisk)end(){
	fmt.Println(f.name,"结束工作，可以弹出")
}
func(f FlashDisk)delete(){
	fmt.Println(f.name,"删除")
}
type Student3 struct{
	Person
	school string
}
type Book struct{
	bookName string
	price float64
}
type Student2 struct{
	name string
	age int
	book *Book
}
type Student struct{
	name string
	age int
	book Book
}
type Worker struct{
	name string
	age int
	sex string
}
type Cat struct{
	color string
	age int
}
func (w Worker)work(){
	fmt.Println(w.name,"在工作")
}
func(p *Worker)rest(){
	fmt.Println(p.name,"在休息")
}
func(p *Worker)printInfo(){
	fmt.Printf("工人姓名：%s,工人年龄：%d,工人性别：%s\n",p.name,p.age,p.sex)
}
func(p *Cat)printInfo(){
	fmt.Printf("猫的颜色：%s,猫的年龄：%d\n",p.color,p.age)
}
func(p *Person)eat(){
	fmt.Println("吃窝窝头")
}
func(s *Student3)study(){
	fmt.Println("学习")
}
func(s *Student3)eat(){
	fmt.Println("吃猫粮")
}
type Person struct{
	name string
	age int
	sex string
	address string
}

func fun2(p1 *int){
	*p1 = 200
}
func fun3()*[4]int{
	arr := [4]int{5,6,7,8}
	fmt.Printf("arr:%p\n",&arr)
	return &arr
}

func increment()func()int{
	i := 0 //返回值为函数，且返回函数用到i，i不会随着外层函数调用销毁，下次调用外层函数时重新创建对应的i
	fun := func()int{
		i++
		return i
	}
	return fun
}

func add(a,b int)int{
	return a+b
}
func sub(a,b int)int{
	return a-b
}
func oper(a,b int,fun func(int,int)int)int{
	fmt.Println(a,b,fun)
	return fun(a,b)
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

// 求1-n的和
func getSum2( n int )(int){
	if( n == 1 ){
		return 1
	}
	sum := getSum2( n - 1 ) + n
	return sum
}

// 斐波那契数列
func getFibonacci(n int)(int){
	if( n == 1|| n == 2 ){
		return 1
	}
	return getFibonacci( n - 1 )+getFibonacci( n - 2 )
}

func fun1(s string){
	fmt.Println(s)
}