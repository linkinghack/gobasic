package main

import "fmt"

func main() {
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

	//slice literals 切片字面量
	q := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(q)

	r := []bool{true, false, true, true, false, true}
	fmt.Println(r)

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
