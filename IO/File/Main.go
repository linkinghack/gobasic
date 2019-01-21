package main

import (
	"fmt"
	"os"
)

func main()  {
	currentFF,err := os.Create("currentFolderFile.txt")
	if(err != nil){
		panic(err)
	}
	defer currentFF.Close()

	dirNames,_ := currentFF.Readdirnames(-1)
	currentFF.WriteString("hello, first file\n")
	fmt.Println(dirNames)

}