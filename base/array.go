package main

import "fmt"

func main() {
	ArrayAsParameter()
	//MultiDimensionalArray()
}

func MultiDimensionalArray() {
	matrix := [3][3]int{{1, 2, 3}, {1, 2, 3}, {1, 2, 3}}
	// cube := [][][]int{}
	for i, v := range matrix {
		// v is an array of in with size 3
		fmt.Printf("Index: %d, Value: %d\n", i, v)
	}
}

func ArrayAsParameter() {
	funcRecvSlice := func(nums []int) []int {
		nums[0] = 100

		return nums
	}

	funcRecvArray := func(nums [5]int) [5]int {
		nums[0] = 1000
		return nums
	}

	funcRecvSlicePointer := func(nums *[]int) {
		*nums = append(*nums, 1000)
	}

	numArray := [...]int{1, 2, 3, 4, 5} // array
	numSlice := []int{1, 2, 3, 4, 5}    // slice

	// 传递切片
	// 切片和数组的外在区别为是否指定长度
	// numArray通过 ... 间接的指定了长度，所以是数组
	// numSlice 则一开始就是一个切片，其对应的数组永远无法获得
	fmt.Printf("Before: %+v\n", numSlice)
	sliceRtn := funcRecvSlice(numSlice[:])
	fmt.Printf("After funcRecvSlice: %+v\n", numSlice) // 原始slice未改变
	fmt.Printf("funcRecvSlice Returned: %+v\n", sliceRtn)
	funcRecvSlicePointer(&numSlice)
	fmt.Printf("After funcRecvSlicePointer: %+v\n", numSlice)

	// 传递数组
	// 数组在函数参数和返回时都是拷贝传递
	// 会发生整个数组的复制
	arrRtn := funcRecvArray(numArray)
	fmt.Printf("After funcRecvArray: %+v\n", numArray)
	fmt.Printf("funcRecvArray return: %+v\n", arrRtn)
	arrRtn[0] = 2000
	fmt.Printf("After edit array returned: origin=%+v, rtn=%+v\n", numArray, arrRtn)

}

func ArraySlice() {
	var a [2]string
	a[0] = "hello"
	a[1] = "world"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)

	//slice 切片
	var s []int = primes[1:4] //s仅是对primes的引用，对s的修改将在primes中表现出来
	s[2] = 100
	fmt.Println("修改后切片:", s)
	fmt.Println("修改后primes:", primes)
}

func SliceLiteral() {
	//slice literals 切片字面量
	q := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(q)

	r := []bool{true, false, true, true, false, true}
	fmt.Println(r)
}

func SliceIndexing() {

	ss := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(ss)

	//切片其他用法
	sli := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13}
	s1 := sli[:]
	s2 := sli[0:]
	s3 := sli[:13]
	//三种相同结果
	fmt.Println(sli, s1, s2, s3)
}
