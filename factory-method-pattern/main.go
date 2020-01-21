package main

import (
	"github.com/NasSilverBullet/design-patterns-in-go/factory-method-pattern/framework"
	"github.com/NasSilverBullet/design-patterns-in-go/factory-method-pattern/idcard"
)

func main() {
	f := framework.NewFactory(idcard.NewFactory())
	c1 := f.Create("結城浩")
	c2 := f.Create("とむら")
	c3 := f.Create("佐藤花子")
	c1.Use()
	c2.Use()
	c3.Use()
}
