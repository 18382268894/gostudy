package main

import (
	"fmt"
	"os"
)

func main(){
	fileName := "src/day10/file/test.txt"   //文件根目录是以GOPATH为基准
	file,err := os.Open(fileName)
	defer file.Close()
	if err != nil{
		fmt.Println("Open the file failed,err:",err)
	}
	var buf = make([]byte,256)
	n,err := file.Read(buf)
	if err != nil {
		fmt.Println("Read the file failed,err:",err)
		return
	}
	fmt.Println(string(buf[:n]))

}