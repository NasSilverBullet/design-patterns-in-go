package main

import "github.com/liquid-dev/design-patterns-in-go/adapter-pattern/printbanner"

func main() {
	p := printbanner.New("Hello")
	p.PrintWeak()
	p.PrintStrong()
}
