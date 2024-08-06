package style

import (
	"io"

	"golazy.dev/lazyml/html"

	"golazy.dev/lazyml"
)

type Style struct {
	Href    string
	Content string
	Media   string
	Data    map[string]string
}

func (s *Style) Element() lazyml.Element {
	opts := []lazyml.Attr{}
	if s.Media != "" {
		opts = append(opts, html.Media(s.Media))
	}

	if s.Data != nil {
		for k, v := range s.Data {
			opts = append(opts, html.DataAttr(k, v))
		}
	}

	if s.Content != "" {
		return html.Style(s.Content, opts)
	}
	return html.Link(html.Href(s.Href), html.Rel("stylesheet"), opts)

}

func (s *Style) WriteTo(w io.Writer) (int64, error) {
	return s.Element().WriteTo(w)
}
