package db

import "fmt"

type InterfaceDB interface {
	MethodA()
	MethodB()
}

type db struct{}

func NewDB() InterfaceDB {
	fmt.Println("實體化 db 物件")
	return &db{}
}

func (thisObj *db) MethodA() {
	fmt.Println("使用 db.MethodA")
}

func (thisObj *db) MethodB() {
	fmt.Println("使用 db.MethodB")
}
