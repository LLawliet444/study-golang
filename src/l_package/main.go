package main

import(
	_ "l_package/pk1"
	// "l_package/utils"
	// "../../../utilstest"
	// "l_package/person"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"database/sql"
)

func main(){
	// utils.Count()

	// p1.MyTest1()

	// var p person.Person = person.Person{"吉良吉影",2,"女"}
	// fmt.Println(p)
	db,err := sql.Open("mysql","root@tcp(127.0.0.1:3306)/my_test?charset=utf8")
	if err != nil{
		fmt.Println("错误信息：",err)
		return
	}
	fmt.Println("连接数据库成功",db)
}