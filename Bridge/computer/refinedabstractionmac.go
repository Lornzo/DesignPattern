package computer

import (
	"fmt"

	"github.com/Lornzo/DesignPattern/Bridge/printer"
)

type RefinedAbstractionMac struct {
	printer printer.IImplementorPrinter
}

func (thisObj *RefinedAbstractionMac) SetPrinter(p printer.IImplementorPrinter) {
	thisObj.printer = p
}

func (thisObj *RefinedAbstractionMac) Print() {
	fmt.Println("Mac print with printer ")
	thisObj.printer.PrintFile()
}
