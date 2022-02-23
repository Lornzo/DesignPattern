package printer

import "fmt"

type ConcreteImplementorHP struct{}

func (thisObj *ConcreteImplementorHP) PrintFile() {
	fmt.Println("HP")
}
