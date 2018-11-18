package main

import "fmt"

var pow = []int{1, 2, 4, 6, 8, 10}

func main() {
	for i, v := range pow {
		fmt.Printf("2**%d = %d \n", i, v)
	}

	// 省略 value提取，仅使用range的index
	var inds = make([]int, 10)

	fmt.Println("indexs only:")
	for i := range inds { //可以给value一个 '_'
		fmt.Printf(" %d ", i)
		inds[i] = i * 3
	}
	fmt.Print("\n")

	// 省略 index， 仅使用value
	for _, val := range inds {
		fmt.Printf(" %d ", val)
	}
	fmt.Print("\n")
}
