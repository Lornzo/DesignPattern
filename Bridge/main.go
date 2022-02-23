package main

import (
	"github.com/Lornzo/DesignPattern/Bridge/computer"
	"github.com/Lornzo/DesignPattern/Bridge/printer"
)

func main() {

	// computers
	var Mac computer.IAbstractionComputer = &computer.RefinedAbstractionWindows{}
	var Windows computer.IAbstractionComputer = &computer.RefinedAbstractionWindows{}
	var Linux computer.IAbstractionComputer = &computer.RefinedAbstractionLinux{}

	// printers
	var Hp printer.IImplementorPrinter = &printer.ConcreteImplementorHP{}
	var Epson printer.IImplementorPrinter = &printer.ConcreteImplementorEpson{}
	var Canon printer.IImplementorPrinter = &printer.ConcreteImplementorCanon{}

	Mac.SetPrinter(Hp)
	Mac.Print()

	Windows.SetPrinter(Epson)
	Windows.Print()

	Linux.SetPrinter(Canon)
	Linux.Print()
}
