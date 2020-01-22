package framework

type Product interface {
	Use(s string)
	CreateClone() Product
}
