package element

import "fmt"

type InterfaceElement interface {
	MethodA()
	MethodB()
}

func NewElement() InterfaceElement {
	fmt.Println("初始化 element物件")
	return &element{}
}

type element struct{}

func (thisObj *element) MethodA() {
	fmt.Println("使用 element.MethodA")
}

func (thisObj *element) MethodB() {
	fmt.Println("使用 element.MethodB")
}
