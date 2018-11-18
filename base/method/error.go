package main

import (
	"fmt"
	"strconv"
)

type Puppy struct {
	name string
}
type ErrPuppy string

func (p ErrPuppy) Error() string { //实现 Error() string 方法以符合error接口
	return "err: " + string(p)
}
func (p Puppy) Bar() error {
	return ErrPuppy(p.name)
}

func main() {
	i, err := strconv.Atoi("2323")
	fmt.Printf("'2323' converted : %d err: %v\n", i, err)

	p := Puppy{"dolingo"}
	if e := p.Bar(); e != nil {
		fmt.Println(e)
	}
}
