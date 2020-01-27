package listfactory

import (
	"fmt"
	"io"

	"github.com/NasSilverBullet/design-patterns-in-go/abstract-factory-pattern/factory"
)

type ListFactory struct{}

func New() *ListFactory {
	return &ListFactory{}
}

func (lf *ListFactory) CreateLink(caption string, url string) factory.Link {
	return &ListLink{
		caption: caption,
		url:     url,
	}
}

func (lf *ListFactory) CreateTray(caption string) factory.Tray {
	return &ListTray{
		caption: caption,
		tray:    make([]factory.Item, 0),
	}
}

func (lf *ListFactory) CreatePage(title string, author string) factory.Page {
	return &ListPage{
		title:   title,
		author:  author,
		content: make([]factory.Item, 0),
	}
}

type ListLink struct {
	caption string
	url     string
}

func (ll *ListLink) MakeHTML() string {
	return fmt.Sprintf("  <li><a href=\"%s\">%s</a></li>\n", ll.url, ll.caption)
}

type ListTray struct {
	caption string
	tray    []factory.Item
}

func (ll *ListTray) Add(i factory.Item) {
	ll.tray = append(ll.tray, i)
}

func (lt *ListTray) MakeHTML() string {
	txt := fmt.Sprintf(`<li>
%s
<ul>
`, lt.caption)

	for _, i := range lt.tray {
		txt += i.MakeHTML()
	}

	txt += fmt.Sprintf(`</ul>
</li>
`)

	return txt
}

type ListPage struct {
	title   string
	author  string
	content []factory.Item
}

func (lp *ListPage) Add(i factory.Item) {
	lp.content = append(lp.content, i)
}

func (lp *ListPage) MakeHTML() string {
	txt := fmt.Sprintf(`<html><head><title>%[1]s</title></head>
<body>
<h1>%[1]s</h1>
<ul>
`, lp.title)

	for _, i := range lp.content {
		txt += i.MakeHTML()
	}

	txt += fmt.Sprintf(`</ul>
<hr><address>%s</address>
</body></html>
`, lp.author)

	return txt
}

func (lp *ListPage) Output(w io.Writer) {
	w.Write([]byte(lp.MakeHTML()))
}
