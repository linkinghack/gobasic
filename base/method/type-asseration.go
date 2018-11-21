package main

import "fmt"

func main() {
	var i interface{} = "hello"
	s, ok := i.(string) // .()
	fmt.Println(s, " is string? ", ok)
	s = "world"
	fmt.Println(i)
	fmt.Println(s == i)
	fmt.Printf("type: i= %T , s=%T \n", i, s)

	// 类型查询 typeswitch
	switch typeofs := i.(type) {
	case int:
		fmt.Println("type is int")
	case string:
		fmt.Println("type is string")
		fmt.Println("typs :", typeofs)
	default:
		fmt.Println("typs :", typeofs)
	}
}
