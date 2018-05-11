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
	fmt.Printf("表达式：%v\n",exps.exp)
}

//将表达式字符串数组转化为字符串
func (exps *Expression) ToString() string {
	var str string = ""
	for i := 0; i < exps.num; i++ {
		str = str + exps.exp[i]
	}
	return str
}

/*func(exps *Expression)Assign(strs []string){
	exps.exp
}*/

/*func Test() {
	var expss = NewExpression()
	expss.add("100")
	expss.add("200")
	fmt.Println(expss.ToString())
}
*/
