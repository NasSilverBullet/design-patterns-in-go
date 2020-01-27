package factory

import "io"

type Factory interface {
	CreateLink(caption string, url string) Link
	CreateTray(caption string) Tray
	CreatePage(title string, author string) Page
}

type factory struct {
	Factory Factory
}

func GetFactory(f Factory) *factory {
	return &factory{
		Factory: f,
	}
}

type Item interface {
	MakeHTML() string
}

type Link interface {
	Item
}

type Tray interface {
	Item
	Add(i Item)
}

type Page interface {
	Add(i Item)
	Output(w io.Writer)
	MakeHTML() string
}
