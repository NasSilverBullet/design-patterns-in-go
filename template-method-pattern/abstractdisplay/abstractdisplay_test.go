package abstractdisplay_test

import (
	"reflect"
	"testing"

	"github.com/NasSilverBullet/design-patterns-in-go/template-method-pattern/abstractdisplay"
)

type mockAbstrctDispaly struct {
	abstractdisplay.AbstractDisplay
	openCount  int
	printCount int
	closeCount int
}

func (m *mockAbstrctDispaly) Open() {
	m.openCount++
}

func (m *mockAbstrctDispaly) Print() {
	m.printCount++
}

func (m *mockAbstrctDispaly) Close() {
	m.closeCount++
}

func TestDisplay_Display(t *testing.T) {
	type fields struct {
		display abstractdisplay.AbstractDisplay
	}
	tests := []struct {
		name   string
		fields fields
		want   *mockAbstrctDispaly
	}{
		{"Success", fields{&mockAbstrctDispaly{}}, &mockAbstrctDispaly{openCount: 1, printCount: 5, closeCount: 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := abstractdisplay.New(tt.fields.display)
			d.Display()
			if got := tt.fields.display; !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Display.Display() Unexpected method call %v, want %v", got, tt.want)
			}
		})
	}
}
