package idcard

import "github.com/NasSilverBullet/design-patterns-in-go/factory-method-pattern/framework"

type IDCardFactory struct {
	owners []string
}

func NewFactory() framework.Factory {
	return &IDCardFactory{}
}

func (icf *IDCardFactory) CreateProduct(o string) framework.Product {
	return New(o)
}

func (icf *IDCardFactory) RegisterProduct(p framework.Product) {
	icf.owners = append(icf.owners, p.(*IDCard).GetOwner())
}

func (icf *IDCardFactory) GetOwners() []string {
	return icf.owners
}
