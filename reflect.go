package main

import (
	"fmt"
	"reflect"
)

type Person struct{
	Name string `json:"name"`
	Age int `json:"age"`
}

// 方法名首字母需大写否则反射跨包检测方法检测不到
func (p Person)Cat(n1 int,n2 int)(a int,b int){
	return n1 + n2,n2
}

func (p Person)Date(){
	fmt.Println(p.Name,p.Age)
}

func Tian(showmaker interface{}){
	rTyp := reflect.TypeOf(showmaker)
	rValue := reflect.ValueOf(showmaker)
	// 获取rvalue.kind
	rValueKind := rValue.Kind()
	// if rvalue!=reflect.value return
	if rValueKind != reflect.Struct{
		return
	}
	// 获取rValue所持有的字段个数
	num := rValue.NumField()

	for i:=0;i<num;i++{
		// tag标签匹配到json则把字段tag返回给tag
		tag := rTyp.Field(i).Tag.Get("json")
		fmt.Printf("value=%v tag=%v",rValue.Field(i),tag)
	}
	// 获取rvalue所持有的方法个数，方法首字母大写方能检测到
	nummuthod := rValue.NumMethod()
	fmt.Println(nummuthod)
	// Method参数为第几个方法，方法排序默认按照函数名排序（ASCII码)
	rValue.Method(1).Call(nil) //call为访问方法的参数,nil为不添加参数
	// 声明一个reflect.Value类型的slice
	var reslice []reflect.Value // 访问方法参数需要为value类型切片
	// 给reslice追加两个元素
	reslice = append(reslice,reflect.ValueOf(20))
	reslice = append(reslice,reflect.ValueOf(30))
	// rValue.Method访问第一个方法，并传入参数reslice
	res := rValue.Method(0).Call(reslice) //传入value类型切片返回value类型切片
	// 由于返回值为value类型切片，方法返回一个值，则取值为res[0]取第一个
	fmt.Println("res=",res[0].Int(),res[1].Int()) 
}

func main(){
	showmaker := Person{"tom",20}
	Tian(showmaker)
}