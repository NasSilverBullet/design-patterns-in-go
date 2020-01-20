package printbanner_test

import (
	"bytes"
	"io"
	"os"
	"reflect"
	"testing"

	"github.com/NasSilverBullet/design-patterns-in-go/adapter-pattern/printbanner"
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

func TestPrintBanner_PrintWeak(t *testing.T) {
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
		t.Run(tt.name, func(t *testing.T) {
			p := printbanner.New(tt.fields.String)
			if got := captureStdout(t, p.PrintWeak); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PrintBanner.PrintWeak() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPrintBanner_PrintStrong(t *testing.T) {
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
		t.Run(tt.name, func(t *testing.T) {
			p := printbanner.New(tt.fields.String)
			if got := captureStdout(t, p.PrintStrong); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PrintBanner.PrintStrong() = %v, want %v", got, tt.want)
			}
		})
	}
}
