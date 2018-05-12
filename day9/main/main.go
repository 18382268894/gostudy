package main

import (
	"day9/postfix"
)

func main() {
	var exps = postfix.NewExpression() //表达式
	var opStack = postfix.NewStack()   //符号栈
	var str string = "(3+4)*5-6"
	postfix.DealStr(str, exps, opStack)
	exps.PrintExps()
}
