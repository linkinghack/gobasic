package main

import (
	"fmt"
	"math"
	"runtime"
	"time"
)

func sqrt(x float64) string {
	if x < 0 { //条件表达式 没有()；  {} 必须
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim { //条件表达式之前执行一个简单的语句, 在此初始化的变量只具有局部作用域(if 和 else块内)
		return v
	}
	return lim
}

func main() {

	//循环控制
	sum := 0
	for i := 0; i < 100; i++ { //注意临时变量i使用:=定义，条件表达没有()，循环体必须有{}
		sum += i
	}
	fmt.Println(sum)

	for sum < 1000 { //初始化和后置语句可选  ;  省略所有条件 for {} 时无限循环
		sum += sum
	}
	fmt.Println(sum)

	//for 当作 while
	for sum < 10000 {
		sum += sum
	}
	fmt.Println(sum)

	// 条件控制
	if 1 > 0; 3 < 2 {
		fmt.Println("multi condition if block")
	} else {
		fmt.Println("multi condition of if successed")
	}

	fmt.Println(sqrt(2), " ", sqrt(-4))
	fmt.Println(pow(3, 9, 10000))

	// 条件选择
	//实际上，Go 自动提供了在这些语言中每个 case 后面所需的 break 语句。
	//除非以 fallthrough 语句结束，否则分支会自动终止。 Go 的另一点重要的不同在于 switch 的 case 无需为常量，且取值不必为整数
	os := runtime.GOOS
	fmt.Println("OS type: ", os)
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.", os)
	}

	fmt.Println(time.Now())

	//无条件switch, 简洁化if-then-else
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good Morning!")
	case t.Hour() < 17:
		fmt.Println("Good Afternoon!")
	default:
		fmt.Println("Good evening!")
	}

	// defer 推迟执行到外层函数返回后。推迟的函数调用，参数会立即求值，但知道外层返回后才调用
	// 多次defer 会压栈， FILO执行
	defer fmt.Println("end")
	fmt.Println("before")

}
