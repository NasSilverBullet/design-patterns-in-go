package messagebox_test

import (
	"bytes"
	"io"
	"os"
	"reflect"
	"testing"

	"github.com/NasSilverBullet/design-patterns-in-go/prototype-pattern/messagebox"
)

func captureStdout(t *testing.T, f func(arg string), arg string) string {
	r, w, err := os.Pipe()
	if err != nil {
		t.Fatal(err)
	}

	stdout := os.Stdout
	os.Stdout = w

	f(arg)

	os.Stdout = stdout
	w.Close()

	var buf bytes.Buffer
	io.Copy(&buf, r)

	return buf.String()
}

func TestMessageBox_Use(t *testing.T) {
	type fields struct {
		decochar string
	}
	type args struct {
		s string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{"Success", fields{decochar: "*"}, args{s: "hoge"}, `********
* hoge *
********
`},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mb := messagebox.New(tt.fields.decochar)
			if got := captureStdout(t, mb.Use, tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MessageBox.Use() print %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMessageBox_CreateClone(t *testing.T) {
	type fields struct {
		decochar string
	}
	tests := []struct {
		name   string
		fields fields
	}{
		{"Success", fields{decochar: "*"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mb := messagebox.New(tt.fields.decochar)
			if got, want := mb.CreateClone(), mb; !reflect.DeepEqual(got, want) {
				t.Errorf("MessageBox.CreateClone() = %v, want %v", got, want)
			}
		})
	}
}
