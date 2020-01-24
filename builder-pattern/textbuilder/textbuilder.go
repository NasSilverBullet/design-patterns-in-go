package textbuilder

import "fmt"

type TextBuilder struct {
	buffer string
}

func New() *TextBuilder {
	return &TextBuilder{}
}

func (tb *TextBuilder) MakeTitle(title string) {
	tb.buffer += "===============================\n"
	tb.buffer += fmt.Sprintf("『%s』\n\n", title)
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

func (tb *TextBuilder) Close() {
	tb.buffer += "===============================\n"
}

func (tb *TextBuilder) GetResult() string {
	return tb.buffer
}
