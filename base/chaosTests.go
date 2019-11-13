package main

import (
	"fmt"
	"runtime"
	"strconv"
	"strings"
)

type TreeNode struct {
	Data  int
	Left  *TreeNode
	Right *TreeNode
}

type Base struct {
	name string
}

func (b Base) hello() {
	fmt.Println(b.name)
}

type Top struct {
	Base // 匿名组合
	addr string
}

func (t Top) where() string {
	return t.addr
}

func main() {
	fmt.Println(runtime.GOOS)

	SliceOfStructEditableTest()

	// t := Top{Base{"linking"}, "Tian Jin"}
	// t.hello()
	// fmt.Println(t.name)

	// sbuilder := strings.Builder{}

	// sbuilder.Write([]byte{68, 78, 98, 112})
	// fmt.Println(sbuilder.String())

	//fmt.Println("default pointer in struct: ",tnode.Right)
	//
	//fmt.Println("count word: ",WordCount("hello world , I am linking"))

	//StringTest()
	//FuncVlaueTest()

	//MapTest()

	//c := 's'
	//fmt.Println(reflect.TypeOf(c))   //int32

	//TwoDimensionalSliceTest()
	// fibonacci test
	/*for i := 1; i < 10; i++ {
		fmt.Println(fib(i))
	}*/
	//defer proc1()

}

func fibonacci() func() int {
	var a, b int
	a = 1
	b = 1
	return func() int {
		t := a
		a = b
		b = b + t
		return t
	}
}

func FuncVlaueTest() {
	fmt.Println(funcvalue(34, 56, func(i int, i2 int) int {
		return i + 2*i2 ^ 2
	}))

}

func funcvalue(x int, y int, fn func(int, int) int) int { // a function parameter
	return fn(x, y)
}

func StringTest() {
	s := "foo,bar;zee"
	fmt.Printf("char at string: %c, type: %T\n", s[2], s[2])
}

func MapTest() {
	var m = map[string]int{}
	m["hello"] = 123
	for i := 0; i < 100; i++ {
		m["h"+strconv.Itoa(i)] = i * 2
	}
	fmt.Println("length of map: ", len(m))
	fmt.Println(m)
}

func Pic(dx, dy int) [][]uint8 {
	var pic = [][]uint8{}
	for i := 0; i < dx; i++ {
		temp := []uint8{}
		for j := 0; j < dy; j++ {
			temp = append(temp, uint8(2*i+j^2))
		}
	}
	return pic
}

func WordCount(s string) map[string]int {
	var words map[string]int = make(map[string]int)
	for _, v := range strings.Fields(s) {
		if _, ok := words[v]; ok {
			words[v]++
		} else {
			words[v] = 1
		}
	}
	return words
}

func TwoDimensionalSliceTest() {
	board := [][]string{}

	for i := 0; i < 10; i++ {
		board = append(board, []string{"hello", "world"})
	}

	fmt.Printf("cap: %d, len: %d \n", cap(board), len(board))
	for j := 0; j < len(board); j++ {
		fmt.Println(board[j])
	}
}

func SliceOfStructEditableTest() {
	type Person struct {
		name string
		age  int
	}

	people := []Person{
		{name: "linking", age: 19},
		{name: "hahah", age: 24},
	}

	for _, p := range people {
		p.name = "changed"
		p.age = 100
	}

	fmt.Printf("People: %+v", people)
}

func SliceTest() {
	arr := [5]int{7, 5, 4, 2, 1}
	var s []int = arr[1:4] // 4 is not included
	fmt.Println("cap: ", cap(s))
	fmt.Println("length: ", len(s))
	s = s[0:cap(s)]
	PrintSlice(&s)

	var arr2 = make([]int, 3, 4)
	fmt.Println("new slice cap: ", cap(arr2))
	//arr2[] = 234
	PrintSlice(&arr2)
}

func PrintSlice(s *[]int) {
	for i, v := range *s {
		fmt.Printf("index: %d, value: %d \n", i, v)
	}
}

func proc1() {
	defer fmt.Println("Hi, this this the 1st process.")
	proc2()
}

func proc2() {
	fmt.Println("And this is the 2nd process.")
	fmt.Println("this is the end of proc2")
}

func fib(n int) int {
	if n < 0 {
		panic("Cannot resolve negative number.")
	}
	if n == 1 || n == 2 {
		return 1
	} else {
		a := fib(n-1) + fib(n-2)
		return a
	}
}
