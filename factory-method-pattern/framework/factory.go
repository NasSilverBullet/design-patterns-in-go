package framework

type Factory interface {
	CreateProduct(owner string) Product
	RegisterProduct(product Product)
}

type factory struct {
	Factory Factory
}

func NewFactory(f Factory) *factory {
	return &factory{
		Factory: f,
	}
}

func (f *factory) Create(owner string) Product {
	p := f.Factory.CreateProduct(owner)
	f.Factory.RegisterProduct(p)
	return p
}
