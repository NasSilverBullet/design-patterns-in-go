package stringdisplay

import (
	"bytes"
	"io"
	"os"
	"testing"
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
		width  int
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sd := &StringDisplay{
				string: tt.fields.string,
				width:  tt.fields.width,
			}
			sd.Open()
		})
	}
}

func TestStringDisplay_Print(t *testing.T) {
	type fields struct {
		string string
		width  int
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sd := &StringDisplay{
				string: tt.fields.string,
				width:  tt.fields.width,
			}
			sd.Print()
		})
	}
}

func TestStringDisplay_Close(t *testing.T) {
	type fields struct {
		string string
		width  int
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sd := &StringDisplay{
				string: tt.fields.string,
				width:  tt.fields.width,
			}
			sd.Close()
		})
	}
}

func TestStringDisplay_PrintLine(t *testing.T) {
	type fields struct {
		string string
		width  int
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sd := &StringDisplay{
				string: tt.fields.string,
				width:  tt.fields.width,
			}
			sd.PrintLine()
		})
	}
}
