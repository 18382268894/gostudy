package postfix

import (
	"errors"
)

type numStack struct {
	elem [256]int
	top  int
}

//创建符号堆栈
func NewNumStack() *numStack {
	st := new(numStack)
	st.top = 0
	return st
}

//压栈
func (st *numStack) pushNum(elem int) error {
	if st.top >= len(st.elem)-1 {
		err := errors.New("Push the element failed,stack is full")
		return err
	}
	st.elem[st.top] = elem
	st.top++
	return nil
}

//出栈
func (st *numStack) popNum() (int, error) {
	var e int
	if st.top == 0 {
		err := errors.New("Pop the element failed,stack is empty")
		return e, err
	}
	e = st.elem[st.top-1]
	st.top--
	return e, nil
}

/*//栈顶元素

func (st *Stack) topNum() int {
	if st.top == 0 {
		return
	}
	return st.elem[st.top-1]
}
*/