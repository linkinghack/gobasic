package main

import "fmt"

type Stuff struct {
	name   string
	gender string
	id     int
}

// 复合字面量
func main() {
	// 使用new 初始化对象
	// 返回的是指针
	p1 := new(Stuff) // var p1 = new(Stuff)
	p1.name = "Stuff 1"
	p1.gender = "female"
	p1.id = 1

	// 使用复合字面量直接初始化
	// 相当于使用new，同样获得了指针
	p2 := &Stuff{"Stuff 2", "male", 2}
	fmt.Println("OBJ1: ", *p1)
	fmt.Println("OBJ2: ", *p2)

}
