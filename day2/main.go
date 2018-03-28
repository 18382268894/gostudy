package main

import (
	"fmt"
)

func main() {
	/*var arr = [2][3]int{
		{1, 2, 3},
		{4, 5, 6},
	}*/
	// rangeArr(arr)
	//copyArr(arr)
	//var arr = [...]int{1, 2, 3, 4, 5}
	//fmt.Printf("arrSum = %d", arrSum(arr))
	var arr = [10]int{0, 1, 3, 8, 4, 13, 5, 2, 19, 7}
	//fmt.Println(sort(arr))
	fmt.Println(sum8(arr))
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

//二维数组是值复制
//复制的副本cp有自己的单独存储空间
//修改副本的值，不影响原arr的值
func copyArr(arr [2][3]int) {
	cp := arr
	cp[0][0] = 100
	fmt.Printf("cp的地址%p：\n", &cp)
	fmt.Printf("arr的地址%p：\n", &arr)
	fmt.Println(cp)
	fmt.Println(arr)
}

//数组求和
func arrSum(arr [5]int) int {
	var sum = 0
	for _, value := range arr {
		sum += value
	}
	return sum
}

//给一维数组元素进行排序(冒泡排序)
func sort(arr [10]int) [10]int {
	length := len(arr)
	for i := 0; i < length; i++ {
		for j := 0; j < length-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	return arr
}

//求数组两个想加等于8的元素的index
func sum8(arr [10]int) [][]int {
	length := len(arr)
	result := [][]int{}
	for i := 0; i < length; i++ {
		for j := i + 1; j < length; j++ {
			if arr[i]+arr[j] == 8 {
				slice := []int{i, j}
				result = append(result, slice)
			}
		}
	}
	return result
}
