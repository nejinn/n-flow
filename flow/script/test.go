package main

import (
	"fmt"
	"nFlow/common/mapFunc"
)

type MyStruct struct {
	i int
	s string
}

func foo0() (interface{}, error) {
	fmt.Println("running foo0: ")
	return []int{100, 200}, nil
}

func foo1(a int) (string, string) {
	fmt.Println("running foo1: ", a)
	return "aaaa", "bbb"
}

func foo2(a, b int, c string) MyStruct {
	fmt.Println("running foo2: ", a, b, c)
	return MyStruct{10, "ccc"}
}

func main() {
	funcs := map[string]interface{}{
		"foo0": foo0,
		"foo1": foo1,
		"foo2": foo2,
	}

	_, _ = mapFunc.CallFcs(funcs["foo0"])

	//var a = "2345"
	//b := strings.Split(a, ",")
	//fmt.Println(b)
}
