package printer

import "fmt"

type ConcreteImplementorEpson struct{}

func (thisObj *ConcreteImplementorEpson) PrintFile() {
	fmt.Println("Epson")
}
