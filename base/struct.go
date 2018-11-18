package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	fmt.Println(Vertex{1, 2})
	v := Vertex{3, 4}
	fmt.Println(v.X) //使用点号访问元素

	// 只想结构体的指针
	p := &v
	fmt.Println((*p).X)
	fmt.Println(p.Y) //允许省略指针取值运算符 - 隐式间接引用

	var (
		v1 = Vertex{2, 3}  //普通结构体
		v2 = Vertex{X: 8}  // name: 方法指定元素值，忽略顺序
		v3 = Vertex{}      //默认值初始化
		po = &Vertex{1, 2} //返回一个指针
	)
	fmt.Println("point: ", v1)
	fmt.Println("point: v2:", v2, " , v3:", v3, " po:", po)

}
