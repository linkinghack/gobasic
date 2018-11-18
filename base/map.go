package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m = map[string]Vertex{
	"Bell Labs": {40.1871, 18.2883},
}

func main() {
	fmt.Println(m)

	// 申请内存
	var mp = make(map[string]int)
	mp["math"] = 100
	mp["english"] = 98
	mp["hahah"] = 2981
	var value int = mp["math"]
	fmt.Println("math : ", value)
	delete(mp, "hahah")                                    //删除map元素
	v, exist := mp["hahah"]                                //接收两个返回值，第一个是值，第二bool代表key是否存在，不存在时，value为map-value的默认0值
	fmt.Println("hahah exits? :", exist, " value is: ", v) //false 0

}
