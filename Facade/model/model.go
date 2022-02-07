package model

import "fmt"

type InterfaceModel interface {
	MethodA()
	MethodB()
}

func NewModel() InterfaceModel {
	fmt.Println("實體化 model 物件")
	return &model{}
}

type model struct{}

func (thisObj *model) MethodA() {
	fmt.Println("使用 model.MethodA")
}

func (thisObj *model) MethodB() {
	fmt.Println("使用 model.MethodB")
}
