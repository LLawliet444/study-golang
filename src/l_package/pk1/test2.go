package pk1

import(
	"l_package/utils"
	"l_package/utils/timeutils"
	"fmt"
)

func MyTest1(){
	utils.Count()
	timeutils.PrintTime()
	fmt.Println("fuck")
}

func init(){
	fmt.Println("pk1包下的init()")
}