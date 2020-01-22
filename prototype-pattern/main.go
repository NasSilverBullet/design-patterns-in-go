package main

import (
	"github.com/NasSilverBullet/design-patterns-in-go/prototype-pattern/framework"
	"github.com/NasSilverBullet/design-patterns-in-go/prototype-pattern/messagebox"
	"github.com/NasSilverBullet/design-patterns-in-go/prototype-pattern/underlinepen"
)

func main() {
	m := framework.NewManager()

	m.Register("strong message", underlinepen.New("~"))
	m.Register("warning box", messagebox.New("*"))
	m.Register("slash box", messagebox.New("/"))

	m.Create("strong message").Use("Hello, world.")
	m.Create("warning box").Use("Hello, world.")
	m.Create("slash box").Use("Hello, world.")
}
