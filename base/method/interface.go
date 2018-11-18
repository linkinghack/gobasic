package main

import (
	"fmt"
	"math"
)

// 定义一个接口，其中包含方法模板
type Abser interface {
	Abs() float64
}

// 不可以给内置对象添加方法，但可以重新定义一个
type MyFloat float64

func (f MyFloat) Abs() float64 { // 实现Abs() 方法后就已经符合Abser接口规范，不需要显示声明
	if f > 0 {
		return float64(f)
	} else {
		return float64(-f)
	}
}

//new type implemented Abser
type Vertex struct {
	X, Y float64
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	var a Abser
	var b = Vertex{3.4, 8.2}
	// a = b // 不可以，因为实现Abs() 的是 *Vertex, Vertex对象本身将不具有Abs()
	a = &b //ok

	fmt.Println(a.Abs())
}
