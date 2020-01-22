package singleton_test

import (
	"testing"

	"github.com/NasSilverBullet/design-patterns-in-go/singleton-pattern/singleton"
)

func TestGetInstance(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"Success"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, want := singleton.GetInstance(), singleton.GetInstance(); got != want {
				t.Errorf("GetInstance() return not same records")
			}
		})
	}
}
