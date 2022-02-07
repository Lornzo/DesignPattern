package main

import (
	"fmt"

	"github.com/Lornzo/DesignPattern/Adapter/adapter"
)

func main() {
	client()
}

func client() {
	var target adapter.InterfaceTarget = adapter.NewTarget()
	var num int64 = 10
	if num > target.GetRandNumber() {
		fmt.Println("num > rand number")
	} else {
		fmt.Println("num <= rand number")
	}
}
