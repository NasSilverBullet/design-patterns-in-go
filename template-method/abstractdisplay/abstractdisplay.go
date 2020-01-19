package abstractdisplay

type AbstractDisplay interface {
	Open()
	Print()
	Close()
}

type Display struct {
	display AbstractDisplay
}

func New(ab AbstractDisplay) *Display {
	return &Display{
		display: ab,
	}
}

func (d *Display) Display() {
	d.display.Open()
	for i := 0; i < 5; i++ {
		d.display.Print()
	}
	d.display.Close()
}
