package htmlbuilder_test

import (
	"bytes"
	"testing"

	"github.com/NasSilverBullet/design-patterns-in-go/builder-pattern/htmlbuilder"
)

func TestHTMLBuilder_MakeTitle(t *testing.T) {
	type args struct {
		title string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"Success", args{title: "hoge"}, "<html><head><title>hoge</title></head></body><h1>hoge</h>"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := bytes.NewBuffer([]byte{})
			hb := htmlbuilder.New(b)
			hb.MakeTitle(tt.args.title)
			if got := string(b.String()); tt.want != got {
				t.Errorf("HTMLBuilder.MakeTitle() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHTMLBuilder_MakeString(t *testing.T) {
	type args struct {
		string string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"Success", args{string: "hoge"}, "<p>hoge</p>"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := bytes.NewBuffer([]byte{})
			hb := htmlbuilder.New(b)
			hb.MakeString(tt.args.string)
			if got := string(b.String()); tt.want != got {
				t.Errorf("HTMLBuilder.MakeString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHTMLBuilder_MakeItems(t *testing.T) {
	type args struct {
		items []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"1Line", args{items: []string{"hoge"}}, "<ul><li>hoge</li></ul>"},
		{"3Line", args{items: []string{"hoge", "huga", "piyo"}}, "<ul><li>hoge</li><li>huga</li><li>piyo</li></ul>"},
		{"0Line", args{items: []string{}}, "<ul></ul>"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := bytes.NewBuffer([]byte{})
			hb := htmlbuilder.New(b)
			hb.MakeItems(tt.args.items)
			if got := string(b.String()); tt.want != got {
				t.Errorf("HTMLBuilder.MakeItems() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHTMLBuilder_Close(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{"1Line", "</body></html>"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := bytes.NewBuffer([]byte{})
			hb := htmlbuilder.New(b)
			hb.Close()
			if got := string(b.String()); tt.want != got {
				t.Errorf("HTMLBuilder.Close() = %v, want %v", got, tt.want)
			}
		})
	}
}
