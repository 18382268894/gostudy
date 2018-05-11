package postfix

import "fmt"

//比较两个运算符的优先级
func CompareOperator(op1 string, op2 string) int {
	if op1 == "+" || op1 == "-" {
		if op2 == "*" || op2 == "/" || op2 == "(" {
			return 1
		}
		if op2 == "+" || op2 == "-" {
			return -1
		}
	}
	if op1 == "*" || op1 == "/" {
		if op2 == "(" {
			return 1
		}else {
			return -1
		}

	}
	if op1 == "(" {
		return 1
	}
	if op1 == "" || op1 == "(" {
		return 1
	}

	return 0
}

//处理输入的运算式
func DealStr(str string, exps *Expression, opStack *Stack) {
	var numStr string = ""

	for i := 0; i < len(str); i++ {
		if str[i] >= 48 && str[i] <= 57 {
			numStr = numStr + string(str[i])
		} else {

			if numStr != "" {
				exps.add(numStr)
				numStr = ""
			}
			if str[i] == ')' {
				dealBracket(opStack, exps)
			} else {
				dealOp(string(str[i]), opStack, exps)
			}

		}
	}
	exps.add(string(str[len(str)-1]))
}

//处理运算符
func dealOp(op2 string, opStack *Stack, exps *Expression) {
	for {
		op1 := opStack.topElem()
		fmt.Printf("栈顶符号：%v\n",op1)
		compResult := CompareOperator(op1, op2)
		if compResult == 1 {
			opStack.push(op2)
			if opStack.top == 0 {
				fmt.Printf("符号栈:%v\n",[]string{""})
			}else {
				fmt.Printf("符号栈:%v\n",opStack.elem[:opStack.top-1])
			}

			break
		}
		if compResult == -1 {
			opStack.pop()
			exps.add(op1)
			fmt.Printf("top：%d\n",opStack.top)
			if opStack.top == 0 {
				fmt.Printf("符号栈:%v\n",[]string{""})
			}else {
				fmt.Printf("符号栈:%v\n",opStack.elem[:opStack.top-1])
			}

		}
	}

}

//处理括号
func dealBracket(opStack *Stack, exps *Expression) {
	for {
		op1 := opStack.topElem()
		fmt.Printf("栈顶符号：%v\n",op1)
		if op1 == "(" {
			opStack.pop()
			break
		}
		opStack.pop()
		exps.add(op1)
	}

}
