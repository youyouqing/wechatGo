package utils

import (
	"fmt"
	"reflect"
)

func AppPrint(i interface{}){
	var kv = make(map[string]interface{})
	vValue := reflect.ValueOf(i)
	vType :=reflect.TypeOf(i)
	for i:=0;i<vValue.NumField();i++{
		kv[vType.Field(i).Name] = vValue.Field(i)
	}
	fmt.Println("获取到数据:")
	for k,v :=range kv{
		fmt.Print(k)
		fmt.Print(":")
		fmt.Print(v)
		fmt.Println()
	}
}
