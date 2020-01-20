package stringdisplay_test

import (
	"bytes"
	"io"
	"os"
	"reflect"
	"testing"

	"github.com/NasSilverBullet/design-patterns-in-go/template-method-pattern/stringdisplay"
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

func TestStringDisplay_Open(t *testing.T) {
	type fields struct {
		string string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{"Success", fields{string: "hoge"}, "+----+\n"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sd := stringdisplay.New(tt.fields.string)
			if got := captureStdout(t, sd.Open); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("StringDisplay.Open() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStringDisplay_Print(t *testing.T) {
	type fields struct {
		string string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{"Success", fields{string: "hoge"}, "|hoge|\n"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sd := stringdisplay.New(tt.fields.string)
			if got := captureStdout(t, sd.Print); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("StringDisplay.Print() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStringDisplay_Close(t *testing.T) {
	type fields struct {
		string string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{"Success", fields{string: "hoge"}, "+----+\n"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sd := stringdisplay.New(tt.fields.string)
			if got := captureStdout(t, sd.Close); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("StringDisplay.Close() = %v, want %v", got, tt.want)
			}
		})
	}
}
