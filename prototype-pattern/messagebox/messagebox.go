package messagebox

import (
	"fmt"
	"strings"

	"github.com/NasSilverBullet/design-patterns-in-go/prototype-pattern/framework"
)

type MessageBox struct {
	decochar string
}

func New(decochar string) *MessageBox {
	return &MessageBox{
		decochar: decochar,
	}
}

func (mb *MessageBox) Use(s string) {
	line := strings.Repeat(mb.decochar, len(s)+4)
	fmt.Printf("%s\n%s %s %s\n%s\n", line, mb.decochar, s, mb.decochar, line)
}

func (mb *MessageBox) CreateClone() framework.Product {
	return &MessageBox{
		decochar: mb.decochar,
	}
}
