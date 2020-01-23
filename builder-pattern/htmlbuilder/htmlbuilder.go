package htmlbuilder

import (
	"fmt"
	"os"
)

type HTMLBuilder struct {
	writer *os.File
}

func New() *HTMLBuilder {
	return &HTMLBuilder{}
}

func (hb *HTMLBuilder) MakeTitle(title string) error {
	f, err := os.OpenFile(fmt.Sprintf("%s.html", title), os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		return err
	}

	hb.writer = f
	hb.writer.Write([]byte(fmt.Sprintf("<html><head><title>%s</title></head></body>", title)))
	hb.writer.Write([]byte(fmt.Sprintf("<h1>%s</h>", title)))

	return nil
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

func (hb *HTMLBuilder) Close() error {
	hb.writer.Write([]byte("</body></html>"))
	if err := hb.writer.Close(); err != nil {
		return err
	}

	return nil
}

func (hb *HTMLBuilder) GetResult() string {
	return hb.writer.Name()
}
