package thirdparty

import (
	"math/rand"
	"strconv"
)

type InterfaceThridParty interface {
	GetRandNumber() (num string)
}

func NewThirdParty() InterfaceThridParty {
	return &thirdparty{}
}

type thirdparty struct{}

func (thisObj *thirdparty) GetRandNumber() (num string) {
	num = strconv.FormatInt(rand.Int63n(100), 10)
	return
}
