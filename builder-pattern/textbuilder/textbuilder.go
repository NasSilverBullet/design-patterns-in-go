package textbuilder

import "fmt"

type TextBuilder struct {
	buffer string
}

func New() *TextBuilder {
	return &TextBuilder{}
}

func (tb *TextBuilder) MakeTitle(title string) error {
	tb.buffer += "===============================\n"
	tb.buffer += fmt.Sprintf("『%s』\n\n", title)
	return nil
}

func (tb *TextBuilder) MakeString(str string) {
	tb.buffer += fmt.Sprintf("■ %s\n\n", str)
}

func (tb *TextBuilder) MakeItems(items []string) {
	for _, item := range items {
		tb.buffer += fmt.Sprintf(" ・%s\n", item)

	}
	tb.buffer += "\n"
}

func (tb *TextBuilder) Close() error {
	tb.buffer += "===============================\n"
	return nil
}

func (tb *TextBuilder) GetResult() string {
	return tb.buffer
}
