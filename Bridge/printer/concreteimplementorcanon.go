package printer

import "fmt"

type ConcreteImplementorCanon struct{}

func (thisObj *ConcreteImplementorCanon) PrintFile() {
	fmt.Println("Canon")
}
