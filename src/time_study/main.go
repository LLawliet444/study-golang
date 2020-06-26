package main

import(
	"time"
	"fmt"
	"math/rand"
)

func main(){
	t1 := time.Now()
	fmt.Printf("%T\n",t1)
	fmt.Println(t1)

	t2 := time.Date(2008,7,15,17,39,20,0,time.Local)
	fmt.Println(t2)

	// 格式化时间
	s1 := t1.Format("2006-01-02 15:04:05")
	fmt.Println(s1)

	// 字符串转time
	s2 := "1999年10月10日"
	t3,err := time.Parse("2006年01月02日",s2)
	if err != nil{
		fmt.Println(err)
		return
	}
	fmt.Println(t3)

	fmt.Println(t1.String())

	y1,m1,day := t2.Date()
	fmt.Println(y1,m1,day)

	hour1,min1,sec1 := t1.Clock()
	fmt.Println(hour1,min1,sec1)

	year2 := t1.Year()
	fmt.Println("年：",year2)
	fmt.Println(t2.YearDay())

	month2 := t1.Month()
	fmt.Println("月：",month2)
	fmt.Println("日",t1.Day())
	fmt.Println("时：",t1.Hour())
	fmt.Println("分：",t1.Minute())
	fmt.Println("秒：",t1.Second())
	fmt.Println("纳秒：",t1.Nanosecond())

	fmt.Println(t1.Weekday())

	// 时间戳
	timeStamp1 := t1.Unix()
	fmt.Println(timeStamp1)

	timeStamp2 := t1.UnixNano()
	fmt.Println(timeStamp2)

	// 时间间隔
	t4 := t1.Add(time.Hour)
	fmt.Println(t1)
	fmt.Println(t4)

	t5 := t1.AddDate(-1,0,0)
	fmt.Println(t5)

	d1 := t4.Sub(t1)
	fmt.Println(d1)
	fmt.Printf("%T\n",d1)

	// 睡眠
	time.Sleep(3*time.Second)
	fmt.Println("mother fuck")

	// 随眠1-10的随机秒
	rand.Seed(time.Now().UnixNano())
	randNum := rand.Intn(10)+1
	fmt.Println(randNum)
	time.Sleep(time.Duration(randNum)*time.Second)
	fmt.Println("mother fuck2")
}