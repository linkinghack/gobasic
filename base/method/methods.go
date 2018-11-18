package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

// 定义一个有 reviver 的function, 则可以再reciver 中指定的type 上使用点号访问此方法
//
func (v Vertex) Abs() float64 {
	return math.Sqrt(math.Pow(v.X, 2) + math.Pow(v.Y, 2))
}

func main() {
	var v = Vertex{3.4, 5.6}
	fmt.Println(v.Abs())
}
