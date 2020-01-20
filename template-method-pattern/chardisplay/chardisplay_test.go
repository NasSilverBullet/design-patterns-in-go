package chardisplay_test

import (
	"bytes"
	"io"
	"os"
	"reflect"
	"testing"

	"github.com/NasSilverBullet/design-patterns-in-go/template-method-pattern/chardisplay"
)

func captureStdout(t *testing.T, f func()) string {
	r, w, err := os.Pipe()
	if err != nil {
		t.Fatal(err)
	}

	stdout := os.Stdout
	os.Stdout = w

	f()

	os.Stdout = stdout
	w.Close()

	var buf bytes.Buffer
	io.Copy(&buf, r)

	return buf.String()
}

func TestCharDisplay_Open(t *testing.T) {
	type fields struct {
		string string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{"Success", fields{string: "hoge"}, "<<"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cd := chardisplay.New(tt.fields.string)
			if got := captureStdout(t, cd.Open); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("StringDisplay.Open() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCharDisplay_Print(t *testing.T) {
	type fields struct {
		string string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{"Success", fields{string: "hoge"}, "hoge"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cd := chardisplay.New(tt.fields.string)
			if got := captureStdout(t, cd.Print); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("StringDisplay.Print() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCharDisplay_Close(t *testing.T) {
	type fields struct {
		string string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{"Success", fields{string: "hoge"}, ">>\n"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cd := chardisplay.New(tt.fields.string)
			if got := captureStdout(t, cd.Close); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("StringDisplay.Close() = %v, want %v", got, tt.want)
			}
		})
	}
}
