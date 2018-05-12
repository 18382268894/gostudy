package postfix

import "fmt"

//表达式
type Expression struct {
	exp [256]string //存放表达式的字符串数组
	num int         //记录存放的个数
}

func NewExpression() *Expression {
	expression := new(Expression)
	expression.num = 0
	return expression
}

//将数字或者运算符加入表达式中
func (exps *Expression) add(str string) {

	exps.exp[exps.num] = str
	exps.num++

}

//输出表达式

func(exps *Expression)PrintExps(){
	fmt.Printf("后缀式表达式：%v\n",exps.exp[:exps.num])
}
