package main

import (
	"fmt"
)

func main() {
	var arr = [2][3]int{
		{1, 2, 3},
		{4, 5, 6},
	}
	rangeArr(arr)
}

//遍历二维数组
func rangeArr(arr [2][3]int) {
	for index, values := range arr {
		for i, v := range values {
			fmt.Printf("arr[%d][%d]=%d\t", index, i, v)
		}
		fmt.Println()
	}
}
