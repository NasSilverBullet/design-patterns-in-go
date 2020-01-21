package idcard_test

import (
	"bytes"
	"io"
	"os"
	"reflect"
	"testing"

	"github.com/NasSilverBullet/design-patterns-in-go/factory-method-pattern/idcard"
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

func TestIDCard_Use(t *testing.T) {
	type fields struct {
		owner string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{"Success", fields{owner: "hoge"}, "hogeのカードを使います。\n"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ic := idcard.New(tt.fields.owner)
			if got := captureStdout(t, ic.Use); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("IDCard.Use() = %v, want %v", got, tt.want)
			}

		})
	}
}

func TestIDCard_GetOwner(t *testing.T) {
	type fields struct {
		owner string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{"Success", fields{owner: "hoge"}, "hoge"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ic := idcard.New(tt.fields.owner)
			if got := ic.GetOwner(); got != tt.want {
				t.Errorf("IDCard.GetOwner() = %v, want %v", got, tt.want)
			}
		})
	}
}
