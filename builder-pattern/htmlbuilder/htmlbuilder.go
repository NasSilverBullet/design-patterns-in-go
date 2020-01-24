package htmlbuilder

import (
	"fmt"
	"io"
)

type HTMLBuilder struct {
	writer io.Writer
}

func New(w io.Writer) *HTMLBuilder {
	return &HTMLBuilder{
		writer: w,
	}
}

func (hb *HTMLBuilder) MakeTitle(title string) {
	hb.writer.Write([]byte(fmt.Sprintf("<html><head><title>%s</title></head></body>", title)))
	hb.writer.Write([]byte(fmt.Sprintf("<h1>%s</h>", title)))
}

func (hb *HTMLBuilder) MakeString(str string) {
	hb.writer.Write([]byte(fmt.Sprintf("<p>%s</p>", str)))
}

func (hb *HTMLBuilder) MakeItems(items []string) {
	hb.writer.Write([]byte("<ul>"))
	for _, item := range items {
		hb.writer.Write([]byte(fmt.Sprintf("<li>%s</li>", item)))
	}
	hb.writer.Write([]byte("</ul>"))
}

func (hb *HTMLBuilder) Close() {
	hb.writer.Write([]byte("</body></html>"))
}
