package banner

import "fmt"

type Banner struct {
	String string
}

func New(s string) *Banner {
	return &Banner{
		String: s,
	}
}

func (b *Banner) ShowWithParen() {
	fmt.Printf("(%s)\n", b.String)
}

func (b *Banner) ShowWithAster() {
	fmt.Printf("*%s*\n", b.String)
}
