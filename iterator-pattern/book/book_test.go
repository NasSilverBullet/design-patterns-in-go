package book

import (
	"testing"
)

func TestBook_GetName(t *testing.T) {
	type fields struct {
		name string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{"Success", fields{name: "hoge"}, "hoge"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := New(tt.fields.name)
			if got := b.GetName(); got != tt.want {
				t.Errorf("Book.GetName() = %v, want %v", got, tt.want)
			}
		})
	}
}
