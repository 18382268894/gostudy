package postfix

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
			if i == len(str)-1{ //处理输入的式子中最后一个数字
				exps.add(numStr)
			}
		} else {
			if numStr != "" {  //将数字存入表达式中
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
	length := opStack.top
	for i:=0;i<length;i++{
		lastOp,_ := opStack.pop()   //处理符号栈中最后的符号

		exps.add(lastOp)
	}

}

//处理运算符
func dealOp(op2 string, opStack *Stack, exps *Expression) {
	for {

		op1 := opStack.topElem()   //获取符号栈顶元素

		compResult := CompareOperator(op1, op2)
		if compResult == 1 {   //如果op2大于栈顶符号，将op2存入符号栈中，并退出当前循环
			opStack.push(op2)
			break
		}
		if compResult == -1 {   //如果op2小于栈顶符号，将op1弹出符号栈并存入表达式中，循环比较op2和下一个栈顶符号
			opStack.pop()
			exps.add(op1)
		}
	}

}

//处理括号
func dealBracket(opStack *Stack, exps *Expression) {
	for {  //弹出右括号前的所有符号存入表达式中，直到遇到左括号，将左括号扔掉
		op1 := opStack.topElem()
		if op1 == "(" {
			opStack.pop()
			break
		}
		opStack.pop()
		exps.add(op1)
	}

}
