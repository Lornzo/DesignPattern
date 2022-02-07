package facade

import (
	"github.com/Lornzo/DesignPattern/Facade/db"
	"github.com/Lornzo/DesignPattern/Facade/element"
	"github.com/Lornzo/DesignPattern/Facade/model"
)

type InterfaceFacadeA interface {
	UseMethodA()
}

type InterfaceFacadeB interface {
	UseMethodB()
}

func NewFacadeA() InterfaceFacadeA {
	return &facade{}
}

func NewFacadeB() InterfaceFacadeB {
	return &facade{}
}

type facade struct {
}

func (thisObj *facade) UseMethodA() {
	var db db.InterfaceDB = db.NewDB()
	var model model.InterfaceModel = model.NewModel()
	var element element.InterfaceElement = element.NewElement()

	db.MethodA()
	model.MethodA()
	element.MethodA()
}

func (thisObj *facade) UseMethodB() {
	var db db.InterfaceDB = db.NewDB()
	var model model.InterfaceModel = model.NewModel()
	var element element.InterfaceElement = element.NewElement()

	db.MethodB()
	model.MethodB()
	element.MethodB()

}
