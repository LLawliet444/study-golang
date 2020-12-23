package main

import (
	"fmt"
	"reflect"
)
func main(){
	// var s float64 = 3.1415
	// v := reflect.ValueOf(s)
	// convertValue := v.Interface().(float64)
	// fmt.Println(convertValue)
	// pointer := reflect.ValueOf(&s)
	// convertPointer := pointer.Interface().(*float64)
	// fmt.Println(*convertPointer)

	p1 := Person{Name:"吉良吉影",Age:2,Sex:"女"}
	// GetMessage(p1)
	CallFunc(p1)
}

type Person struct {
	Name string
	Age int
	Sex string
}

func (p Person) Say (msg string){
	fmt.Println("hello,",msg)
}

func (p Person) Test (i,j int,msg string){
	fmt.Println(i,j,msg)
}

func (p Person) PrintInfo(){
	fmt.Printf("姓名：%s,年龄：%d，性别：%s\n",p.Name,p.Age,p.Sex)
}

func GetMessage( input interface{} ){
	getType := reflect.TypeOf(input)
	getValue := reflect.ValueOf(input)
	// fmt.Println("get Type is:",getType.Name())
	// fmt.Println("get kind is:",getType.Kind())
	fmt.Println(getType)
	fmt.Println("get value is:",getValue)

	// 获取字段
	for i := 0;i < getType.NumField();i++{
		field := getType.Field(i)
		value := getValue.Field(i)
		fmt.Println("field:",field.Name,field.Type)
		fmt.Println(value.Interface())
	}

	// 获取方法
	for i:= 0;i < getType.NumMethod();i++{
		method := getType.Method(i)
		fmt.Println(method)
	}
}

func CallFunc(input interface{}){
	getValue := reflect.ValueOf(input)
	method1 := getValue.MethodByName("PrintInfo")
	method1.Call(nil)
	arg1 := make([]reflect.Value,0)
	method1.Call(arg1)

	method2 := getValue.MethodByName("Say")
	arg2 := []reflect.Value{reflect.ValueOf("反射机制")}
	method2.Call(arg2)

	method3 := getValue.MethodByName("Test")
	arg3 := []reflect.Value{reflect.ValueOf(100),reflect.ValueOf(200),reflect.ValueOf("hello")}
	method3.Call(arg3)
}