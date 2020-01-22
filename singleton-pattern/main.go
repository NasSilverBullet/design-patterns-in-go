package main

import (
	"fmt"

	"github.com/NasSilverBullet/design-patterns-in-go/singleton-pattern/singleton"
)

func main() {
	s1 := singleton.GetInstance()
	s2 := singleton.GetInstance()
	if s1 != s2 {
		fmt.Println("s1とs2は違う")
		return
	}
	fmt.Println("s1とs2は同じ")
}
