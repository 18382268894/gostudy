package main

import (
	"os"
	"fmt"
	"bufio"
	"io"
)

const (
	OPENERR int = -1
	READERR  = -2
)

func main(){
	fileName := "1.txt"
	file,err := os.Open(fileName)
	defer file.Close()
	if err != nil{
		fmt.Println("Open file faild,err:",err)
		os.Exit(OPENERR)
	}
	reader := bufio.NewReader(file)
	for{

		line,err := reader.ReadString('\n')
		if err == io.EOF {
			fmt.Println("Read file finished")
			break
		}
		if err != nil{
			fmt.Println("Read file faild,err:",err)
			os.Exit(READERR)
		}
		fmt.Print(line)
	}
}