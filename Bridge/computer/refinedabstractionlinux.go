package computer

import (
	"fmt"

	"github.com/Lornzo/DesignPattern/Bridge/printer"
)

type RefinedAbstractionLinux struct {
	printer printer.IImplementorPrinter
}

func (thisObj *RefinedAbstractionLinux) SetPrinter(p printer.IImplementorPrinter) {
	thisObj.printer = p
}

func (thisObj *RefinedAbstractionLinux) Print() {
	fmt.Println("Linux print with printer ")
	thisObj.printer.PrintFile()
}
