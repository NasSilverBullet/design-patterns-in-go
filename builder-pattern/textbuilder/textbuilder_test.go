package textbuilder_test

import (
	"testing"

	"github.com/NasSilverBullet/design-patterns-in-go/builder-pattern/textbuilder"
)

func TestTextBuilder_MakeTitle(t *testing.T) {
	type args struct {
		title string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"Success", args{title: "hoge"}, "===============================\n『hoge』\n\n"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tb := textbuilder.New()
			tb.MakeTitle(tt.args.title)
			if got := tb.GetResult(); tt.want != got {
				t.Errorf("TextBuilder.MakeTitle() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTextBuilder_MakeString(t *testing.T) {
	type args struct {
		title string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"Success", args{title: "hoge"}, "■ hoge\n\n"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tb := textbuilder.New()
			tb.MakeString(tt.args.title)
			if got := tb.GetResult(); tt.want != got {
				t.Errorf("TextBuilder.MakeString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTextBuilder_MakeItems(t *testing.T) {
	type args struct {
		items []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"1Line", args{items: []string{"hoge"}}, " ・hoge\n\n"},
		{"3Line", args{items: []string{"hoge", "huga", "piyo"}}, " ・hoge\n ・huga\n ・piyo\n\n"},
		{"0Line", args{items: []string{}}, "\n"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tb := textbuilder.New()
			tb.MakeItems(tt.args.items)
			if got := tb.GetResult(); tt.want != got {
				t.Errorf("TextBuilder.MakeItems() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTextBuilder_Close(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{"Success", "===============================\n"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tb := textbuilder.New()
			tb.Close()
			if got := tb.GetResult(); tt.want != got {
				t.Errorf("TextBuilder.MakeClose() = %v, want %v", got, tt.want)
			}
		})
	}
}
