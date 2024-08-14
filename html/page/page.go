// package document provides helpers to generate an html document
package page

import (
	"io"
	"strings"

	"golazy.dev/lazyml"

	. "golazy.dev/lazyml/html"
	"golazy.dev/lazyml/html/script"
	"golazy.dev/lazyml/html/style"

	"golazy.dev/lazyassets"
)

type Page struct {
	Assets      *lazyassets.Storage
	Styles      []style.Style
	Scripts     []script.Script
	Description string
	Keywords    string
	Charset     string
	Head        []io.WriterTo
	Lang        string
	Title       string
	Viewport    string
	Content     any
	//ImportMap   lazycomponent.ImportMap
	//Components  []lazycomponent.Component
}

func (p *Page) AddScript(s any) {
	p.Scripts = append(p.Scripts, script.New(s))
}

func (p *Page) AddHead(h ...io.WriterTo) *Page {
	p.Head = append(p.Head, h...)
	return p
}

//func (p *Page) Use(c lazycomponent.Component) *Page {
//	p.Components = append(p.Components, c)
//	return p
//}

func (pold Page) Add(args ...any) *Page {
	p := pold
	for _, a := range args {
		switch a := a.(type) {
		//		case lazycomponent.Component:
		//			p.Components = append(p.Components, a)
		case style.Style:
			p.Styles = append(p.Styles, a)
		case script.Script:
			p.Scripts = append(p.Scripts, a)
		case io.WriterTo:
			p.Head = append(p.Head, a)
		default:
			panic("unknown type")
		}
	}
	return &p
}

//func (p *Page) AddStylesheet(ss *lazyassets.Stylesheet) *Page {
//	p.AddStyleLink(ss.Path)
//	return p
//}

func (p *Page) AddStyleLink(href string) *Page {
	if len(href) == 0 {
		panic("href must not be empty")
	}
	if href[0] != '/' {
		href = "/" + href
	}

	if p.Assets != nil {
		//href = p.Assets.Get(href)
	}

	s := style.Style{
		Href: href,
		Data: map[string]string{"turbo-reload": "true"},
	}
	p.Styles = append(p.Styles, s)
	return p
}

func (p *Page) AddStyle(s style.Style) *Page {
	p.Styles = append(p.Styles, s)
	return p
}

func (p *Page) Element() *lazyml.Element {
	var lang io.WriterTo
	if p.Lang != "" {
		lang = Lang(p.Lang)
	}
	var title io.WriterTo
	if p.Title != "" {
		title = Title(p.Title)
	}
	var viewport io.WriterTo
	if p.Viewport != "" {
		viewport = Meta(Name("viewport"), ContentAttr(p.Viewport))
	}
	var charset io.WriterTo
	if p.Charset != "" {
		charset = Meta(Charset(p.Charset))
	}

	var description io.WriterTo
	if p.Description != "" {
		description = Meta(Name("description"), ContentAttr(p.Description))
	}

	var keywords io.WriterTo
	if p.Keywords != "" {
		keywords = Meta(Name("keywords"), ContentAttr(p.Keywords))
	}

	e := Html(
		lang,
		Head(
			charset,
			viewport,
			title,
			description,
			keywords,
			p.head(),
			p.styles(),
			p.importMaps(),
			p.scripts(),
		),
		Body(p.Content),
	)

	return &e

}

func (p *Page) With(content ...any) io.WriterTo {
	p.Content = content
	return p.Element()
}

func (p *Page) WriteTo(w io.Writer) (int64, error) {
	return p.Element().WriteTo(w)
}

func (p *Page) head() []io.WriterTo {

	head := []io.WriterTo{}
	//	for _, c := range p.Components {
	//		if c, ok := (c).(lazycomponent.ComponentWithHead); ok {
	//			head = append(head, c.PageHead()...)
	//		}
	//	}
	head = append(head, p.Head...)

	return head
}

func (p *Page) scripts() []io.WriterTo {
	scripts := []io.WriterTo{}
	//	for _, c := range p.Components {
	//		if c, ok := (c).(lazycomponent.ComponentWithScripts); ok {
	//			for _, s := range c.PageScripts() {
	//
	//				if s.Src != "" {
	//					p, f := p.Assets.Permalink(s.Src)
	//					if f == nil {
	//						panic("File not found: " + s.Src)
	//					}
	//					s.Src = p
	//					s.Integrity = f.Integrity()
	//				}
	//
	//				scripts = append(scripts, s.Element())
	//			}
	//		}
	//	}
	for _, s := range p.Scripts {
		if p.Assets != nil {
			f := p.Assets.Find(s.Src)
			if f == nil {
				panic("File not found: " + s.Src)
			}
			s.Src = f.Permalink()
			s.Integrity = f.Integrity()
		}
		scripts = append(scripts, s.Element())
	}

	return scripts

}

func (p *Page) toPermalink(path string) string {
	if p.Assets == nil {
		return path
	}
	if strings.Contains(path, "://") || p.Assets == nil {
		return path
	}
	if len(path) < 1 {
		return path
	}

	f := p.Assets.Find(path[1:])
	if f == nil {
		panic("File not found:" + path)
	}
	return f.Permalink()
}

func (p *Page) findAsset(asset_path string) lazyassets.File {
	if p.Assets == nil {
		return nil
	}
	f := p.Assets.Find(asset_path)
	return f
}

func (p *Page) importMaps() io.WriterTo {

	//	im := make(lazycomponent.ImportMap)
	//
	//	for _, c := range p.Components {
	//		if c, ok := (c).(lazycomponent.ComponentWithMaps); ok {
	//			im.Merge(c.ImportMap())
	//		}
	//	}
	//
	//	im.Merge(p.ImportMap)
	//
	//	if len(im) == 0 {
	//		return nil
	//	}
	//
	//	// Use permalinks
	//	if p.Assets != nil {
	//		for k, v := range im {
	//			im[k] = p.toPermalink(v)
	//		}
	//	}
	//	mapJSON, err := json.Marshal(struct {
	//		ImportMap map[string]string `json:"imports"`
	//	}{im})
	//
	//	if err != nil {
	//		panic(err)
	//	}
	//
	// return Script(Type("importmap"), mapJSON)

	return Script(Type("importmap"), "{}")

}

func (p *Page) styles() []io.WriterTo {
	// TODO: Get Component Styles
	styles := []io.WriterTo{}
	//	for _, c := range p.Components {
	//		if c, ok := (c).(lazycomponent.ComponentWithStyles); ok {
	//			for _, s := range c.PageStyles() {
	//				if s.Href == "" {
	//					styles = append(styles, s.Element())
	//					continue
	//				}
	//				href := p.toPermalink(s.Href)
	//				styles = append(styles,
	//					&style.Style{
	//						Href:    href,
	//						Content: s.Content,
	//						Media:   s.Media,
	//					})
	//			}
	//		}
	//	}
	for _, s := range p.Styles {
		if s.Href != "" {
			s.Href = p.toPermalink(s.Href)
		}
		styles = append(styles, s.Element())
	}
	return styles
}

func (p *Page) String() string {
	return p.Element().String()
}
