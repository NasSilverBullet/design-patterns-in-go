package underlinepen

import (
	"fmt"
	"strings"

	"github.com/NasSilverBullet/design-patterns-in-go/prototype-pattern/framework"
)

type UnderlinePen struct {
	ulchar string
}

func New(s string) *UnderlinePen {
	return &UnderlinePen{
		ulchar: s,
	}
}

func (up *UnderlinePen) Use(s string) {
	line := strings.Repeat(up.ulchar, len(s))
	fmt.Printf("\" %s \"\n  %s\n", s, line)
}

func (up *UnderlinePen) CreateClone() framework.Product {
	return &UnderlinePen{
		ulchar: up.ulchar,
	}
}
