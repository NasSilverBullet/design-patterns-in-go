package director

import "github.com/NasSilverBullet/design-patterns-in-go/builder-pattern/builder"

type Director struct {
	builder builder.Builder
}

func New(b builder.Builder) *Director {
	return &Director{
		builder: b,
	}
}

func (d *Director) Construct() error {
	if err := d.builder.MakeTitle("Greeting"); err != nil {
		return err
	}

	d.builder.MakeString("朝から昼にかけて")
	d.builder.MakeItems([]string{
		"おはようございます。",
		"こんにちは。",
	})

	d.builder.MakeString("夜に")
	d.builder.MakeItems([]string{
		"こんばんは。",
		"おやすみなさい。",
		"さようなら",
	})

	if err := d.builder.Close(); err != nil {
		return err
	}
	return nil
}
