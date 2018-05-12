package postfix

import (
	"strconv"
)

func(exps *Expression)Calc()int{   //将后缀式计算为数值
	numStack := NewNumStack()
	for i:=0;i<exps.num;i++ {
		num,err := strconv.Atoi(exps.exp[i])
		if err != nil {
			num2,_ := numStack.popNum()
			num1,_ := numStack.popNum()
			var mu int
			switch exps.exp[i] {
			case "*":
				mu= num1 * num2
			case "+":
				mu = num1 + num2
			case "-":
				mu = num1 - num2
			case "/":
				mu = num1 / num2
			}
			numStack.pushNum(mu)

		}else{
			numStack.pushNum(num)
		}
	}
	return numStack.elem[0]
}