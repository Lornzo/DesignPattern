package computer

import (
	"fmt"

	"github.com/Lornzo/DesignPattern/Bridge/printer"
)

type RefinedAbstractionWindows struct {
	printer printer.IImplementorPrinter
}

func (thisObj *RefinedAbstractionWindows) SetPrinter(p printer.IImplementorPrinter) {
	thisObj.printer = p
}

func (thisObj *RefinedAbstractionWindows) Print() {
	fmt.Println("Windows print with printer ")
	thisObj.printer.PrintFile()
}
