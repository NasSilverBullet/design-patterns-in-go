package banner_test

import (
	"bytes"
	"io"
	"os"
	"reflect"
	"testing"

	"github.com/liquid-dev/design-patterns-in-go/adapter-pattern/banner"
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

func TestBanner_ShowWithParen(t *testing.T) {
	type fields struct {
		String string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{"Success", fields{String: "hoge"}, "(hoge)\n"},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			b := banner.New(tt.fields.String)
			if got := captureStdout(t, b.ShowWithParen); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Banner.ShowWithParen() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBanner_ShowWithAster(t *testing.T) {
	type fields struct {
		String string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{"Success", fields{String: "hoge"}, "*hoge*\n"},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			b := banner.New(tt.fields.String)
			if got := captureStdout(t, b.ShowWithAster); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Banner.ShowWithParen() = %v, want %v", got, tt.want)
			}
		})
	}
}
