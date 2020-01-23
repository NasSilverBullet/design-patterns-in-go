package director_test

import (
	"strings"
	"testing"

	"github.com/NasSilverBullet/design-patterns-in-go/builder-pattern/builder"
	"github.com/NasSilverBullet/design-patterns-in-go/builder-pattern/director"
)

type mockBuilder struct {
	builder.Builder
	text string
}

func (m *mockBuilder) MakeTitle(title string) error {
	m.text += title
	return nil
}

func (m *mockBuilder) MakeString(str string) {
	m.text += str
}

func (m *mockBuilder) MakeItems(items []string) {
	m.text += strings.Join(items, "")
}

func (m *mockBuilder) Close() error {
	m.text += "おわり"
	return nil
}

func TestDirector_Construct(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{"Success", "Greeting朝から昼にかけておはようございます。こんにちは。夜にこんばんは。おやすみなさい。さようならおわり"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &mockBuilder{}
			d := director.New(m)
			d.Construct()
			if got := m.text; got != tt.want {
				t.Errorf("Director.Construct() = %v, want %v", got, tt.want)
			}
		})
	}
}
