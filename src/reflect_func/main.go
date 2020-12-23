package main

import (
	"fmt"
	"reflect"
	"strconv"
)
func main(){

	v1 := reflect.ValueOf(fun1)
	v2 := reflect.ValueOf(fun2)
	v3 := reflect.ValueOf(fun3)
	// fmt.Println(v1.Kind(),v1.Type())
	// fmt.Println(v2.Kind(),v2.Type())
	// fmt.Println(v3.Kind(),v3.Type())

	v1.Call(nil)
	v2.Call([]reflect.Value{reflect.ValueOf(100),reflect.ValueOf("hello")})
	s := v3.Call([]reflect.Value{reflect.ValueOf(100),reflect.ValueOf("hello")})
	fmt.Println(s[0].Kind(),s[0].Type(),s[0].Value)
}

func fun1(){
	fmt.Println("fun1无参数")
}

func fun2(i int,s string){
	fmt.Println("fun2参数",i,s)
}

func fun3(i int,s string)(string){
	fmt.Println("fun3参数",i,s,"也有返回值")
	return s + strconv.Itoa(i)
}
