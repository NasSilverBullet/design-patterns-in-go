package printbanner

import (
	"github.com/liquid-dev/design-patterns-in-go/adapter-pattern/banner"
	"github.com/liquid-dev/design-patterns-in-go/adapter-pattern/print"
)

type PrintBanner struct {
	Banner *banner.Banner
}

func New(s string) print.Print {
	return &PrintBanner{
		Banner: banner.New(s),
	}
}

func (pb *PrintBanner) PrintWeak() {
	pb.Banner.ShowWithParen()
}

func (pb *PrintBanner) PrintStrong() {
	pb.Banner.ShowWithAster()
}
