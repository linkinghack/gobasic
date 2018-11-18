package main

import "fmt"

func main() {
	var i interface{} = "hello"
	s := i.(string)
	fmt.Println(s)
	s = "world"
	fmt.Println(i)
	fmt.Println(s==i)
	fmt.Printf("type: i= %T , s=%T \n",i,s)
}

