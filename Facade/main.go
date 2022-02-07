package main

import "github.com/Lornzo/DesignPattern/Facade/facade"

func main() {
	clientA()
	clientB()
}

func clientA() {
	var facade facade.InterfaceFacadeA = facade.NewFacadeA()
	facade.UseMethodA()
}

func clientB() {
	var facade facade.InterfaceFacadeB = facade.NewFacadeB()
	facade.UseMethodB()
}
