package adapter

import (
	"strconv"

	"github.com/Lornzo/DesignPattern/Adapter/thirdparty"
)

type InterfaceTarget interface {
	GetRandNumber() (num int64)
}

func NewTarget() InterfaceTarget {
	return &adapter{}
}

type adapter struct{}

func (thisObj *adapter) GetRandNumber() (num int64) {
	var third thirdparty.InterfaceThridParty = thirdparty.NewThirdParty()
	num, _ = strconv.ParseInt(third.GetRandNumber(), 10, 64)
	return
}
