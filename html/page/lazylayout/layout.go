package lazylayout

import (
	_ "embed"
	"io"

	. "golazy.dev/lazyml/html"
	"golazy.dev/lazyml/html/page"
	"golazy.dev/lazyml/html/script"
	"golazy.dev/lazyml/html/style"
)

//go:embed style.css
var css string

var Layout = &page.Page{
	Lang:     "en",
	Title:    "golazy",
	Viewport: "width=device-width",
	Styles:   []style.Style{{Content: css}},
	Scripts: []script.Script{
		{Src: "https://cdn.skypack.dev/@hotwired/turbo"},
	},
}

func PageHeader() io.WriterTo {
	return Header(H1("lazygo"))
}

func PageNav() io.WriterTo {
	return Nav(
		Ul(
			Li(A(Href("#"), "Web")),
			Li(A(Href("#"), "Docs")),
			Li(A(Href("#"), "Code")),
			Li(A(Href("#"), "Config")),
		),
	)
}
