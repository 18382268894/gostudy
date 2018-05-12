package main

import (
	"day9/postfix"
	"fmt"

	"github.com/urfave/cli"
	"os"
)

func main() {


	app := cli.NewApp()
	app.Name = "Calculator"
	app.Usage = "Calc expression"

	app.Action = func(c *cli.Context) error{
		calc()
		return nil
	}
	app.Run(os.Args)

}

func calc(){
	for{
		fmt.Printf("请输入你的需要计算的式子：\n")
		var str string
		fmt.Scanf("%s",&str)
		var exps = postfix.NewExpression() //表达式
		var opStack = postfix.NewStack()   //符号栈
		postfix.DealStr(str, exps, opStack)
		exps.PrintExps()
		fmt.Printf("计算结果：%d\n",exps.Calc())
	}

}
