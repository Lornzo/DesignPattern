package computer

import "github.com/Lornzo/DesignPattern/Bridge/printer"

type IAbstractionComputer interface {
	SetPrinter(printer.IImplementorPrinter)
	Print()
}
