package main

import "golang.org/x/tour/tree"
import "fmt"
//import "sync"

//var wg  sync.WaitGroup
type NodeStack [](*tree.Tree)
func (h *NodeStack) Len() int { return len(*h)}
func (h *NodeStack) Push(e *tree.Tree){
	*h = append(*h,e)
}
func (h *NodeStack) Top() *tree.Tree {
	return (*h)[h.Len()-1]
}
func (h *NodeStack)Pop() *tree.Tree {
	old := *h
	size := old.Len()
	x := old[size-1]
	*h = old[0:size-1]
	return x
}

func Walk(t *tree.Tree, ch chan int){
	//defer wg.Done()
	var stack NodeStack
	//DFS
	stack.Push(t)
	// 走到最左
	for t!=nil || stack.Len() > 0 {
		if t != nil {  // 走到左子树最低端
			stack.Push(t)
			t = t.Left
		} else {
			t = stack.Pop()
			ch <- t.Value  // 访问结点
			t = t.Right // 访问右子树
		}
	}
	//fmt.Println("walk done")
	//fmt.Println("len of c: ",len(ch))
	close(ch)
}

func Same(t1, t2 *tree.Tree) bool {
	c1 := make(chan int,15)
	c2 := make(chan int,15)
	//wg.Add(2)
	go Walk(t1, c1)
	go Walk(t2, c2)
	//wg.Wait()

	for i:=10; i>0; i-- {
		v1, ok1 := <-c1
		v2, ok2 := <-c2
		//fmt.Println("t1: ",v1, " t2: ",v2, " f1,f2: ",ok1," ",ok2)
		if ok1 && ok2{
			if v1 != v2{
				return false
			}
		}
	}
/*	_, ok1 := <-c1
	_, ok2 := <-c2
	if ok1 || ok2{
		return false
	}*/
	return true
}

func main() {
	fmt.Println(Same(tree.New(1), tree.New(2)))
}