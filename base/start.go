package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
) //声明多个变量

// bool

// string

// int  int8  int16  int32  int64
// uint uint8 uint16 uint32 uint64 uintptr

// byte // uint8 的别名

// rune // int32 的别名
//     // 表示一个 Unicode 码点

// float32 float64

// complex64 complex128

func main() {
	fmt.Println("Hello golang pig!")
	fmt.Printf("Now you have %g problems.", math.Sqrt(7))
	fmt.Printf("Pi: %f \n", math.Pi)

	fmt.Println("1024+2048=", add(1024, 2048))
	a, b := swap("hello", "world")
	fmt.Println("swap: ", a, b)
	math.Abs
	x, y := split(30)
	fmt.Println("命名返回值 Split(sum) ", x, y)

	// 使用var定义变量, 在最后指定类型
	var c, python, java bool = 1, true, false, "no!" //为多个变量赋初值
	var numb = 2                                     //指定初始值可以省略类型
	varible := "hello"                               // := 的方式在类型明确的地方代替var声明（仅在函数内可以）

	fmt.Println(c, python, java, numb, varible)

	fmt.Println("pi: ", Pi)

	var char = 'h' //字符将直接使用int32存储
	fmt.Printf("c is of type %T, value in ASCII: %d\n", char, char)

	fmt.Println(1 / 2) //=0
}

func add(x int, y int) int {
	return x + y
}

// 多值返回, 同类型参数可只写最后一个参数类型
func swap(x, y string) (string, string) {
	return y, x
}

// 命名返回值
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	//只写return 直接返回已命名的返回值
	return
}
