package idcard

import "fmt"

type IDCard struct {
	owner string
}

func New(o string) *IDCard {
	fmt.Printf("%sのカードを作ります。\n", o)
	return &IDCard{
		owner: o,
	}
}

func (ic *IDCard) Use() {
	fmt.Printf("%sのカードを使います。\n", ic.owner)
}

func (ic *IDCard) GetOwner() string {
	return ic.owner
}
