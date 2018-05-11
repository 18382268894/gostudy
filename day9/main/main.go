package main

import (
	"fmt"

	"day9/postfix"
)

func main() {
	var exps = postfix.NewExpression() //表达式
	var opStack = postfix.NewStack()   //符号栈
	var str string = "21*(93+6/3-51)+4-19"
	postfix.DealStr(str, exps, opStack)
	expString := exps.ToString()
	fmt.Printf("表达式：%v",exps)
    fmt.Printf("%s",expString)
}
