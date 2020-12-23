package main

import (
	"fmt"
	"reflect"
)
func main(){
	// var num float64 = 3.1415926

	// pointer := reflect.ValueOf(&num)
	// newValue := pointer.Elem()
	// fmt.Println(newValue.Type(),newValue.CanSet())
	// newValue.SetFloat(5.5)
	// fmt.Println(num)

	cat := Cat{Name:"林克",Age:1,Color:"blue"}
	pointer := reflect.ValueOf(&cat)
	if pointer.Kind() == reflect.Ptr{
		value := pointer.Elem()
		fmt.Println(value.CanSet())
		f1 := value.FieldByName("Name")
		f1.SetString("吉良吉影")
		f3 := value.FieldByName("Color")
		f3.SetString("tiger pattern")
		fmt.Println(cat)
	}
}

type Cat struct{
	Name string
	Age int
	Color string
}

