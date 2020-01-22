package underlinepen_test

import (
	"bytes"
	"io"
	"os"
	"reflect"
	"testing"

	"github.com/NasSilverBullet/design-patterns-in-go/prototype-pattern/underlinepen"
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

func TestUnderlinePen_Use(t *testing.T) {
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
		{"Success", fields{decochar: "~"}, args{s: "hoge"}, `" hoge "
  ~~~~
`},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			up := underlinepen.New(tt.fields.decochar)
			if got := captureStdout(t, up.Use, tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UnderlinePen.Use() print \n%v, want \n%v", got, tt.want)
			}
		})
	}
}

func TestUnderlinePen_CreateClone(t *testing.T) {
	type fields struct {
		ulchar string
	}
	tests := []struct {
		name   string
		fields fields
	}{
		{"Success", fields{ulchar: "~"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			up := underlinepen.New(tt.fields.ulchar)
			if got, want := up.CreateClone(), up; !reflect.DeepEqual(got, want) {
				t.Errorf("UnderlinePen.CreateClone() = %v, want %v", got, want)
			}
		})
	}
}
