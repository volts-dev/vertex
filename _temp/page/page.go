package page

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/volts-dev/vertex/core/window"

	"github.com/volts-dev/vertex/core/html"
	"github.com/volts-dev/vertex/core/js"
)

type (
	Preload struct {
		Type          string
		As            string
		Href          string
		FetchPriority string
	}

	___requestPage struct {
		url        *url.URL
		resolveURL func(string) string

		title          string
		lang           string
		description    string
		author         string
		keywords       string
		preloads       []Preload
		loadingLabel   string
		image          string
		width          int
		height         int
		twitterCardMap map[string]string
	}

	Page struct {
		url            *url.URL
		resolveURL     func(string) string
		title          string
		lang           string
		description    string
		author         string
		keywords       string
		preloads       []Preload
		loadingLabel   string
		image          string
		width          int
		height         int
		twitterCardMap map[string]string
	}
)

func ____makeRequestPage(origin *url.URL, resolveURL func(string) string) Page {
	return Page{
		url:        origin,
		resolveURL: resolveURL,
	}
}

func newPage(origin *url.URL, resolveURL func(string) string) Page {
	return Page{url: origin, resolveURL: resolveURL}
}

func (p Page) Title() string {
	return window.DefaultWindow().Document().Title()
}

func (p Page) SetTitle(format string, v ...any) {
	title := fmt.Sprintf(format, v...)
	window.DefaultWindow().Get("document").Set("title", title)
	if ele, err := html.ToElement(p.metaByProperty("og:title")); err == nil {
		ele.SetAttribute("content", title)
	}
}

func (p Page) Lang() string {
	return js.ValueToString(window.DefaultWindow().
		Get("document").
		Get("documentElement").
		Get("lang"))
}

func (p Page) SetLang(v string) {
	window.DefaultWindow().
		Get("document").
		Get("documentElement").
		Set("lang", v)
}
func (p Page) metaByName(v string) html.IHTMLElement {
	metaValue := window.DefaultWindow().
		Get("document").
		Call("querySelector", "meta[name='"+v+"']")

	if metaValue.IsNull() {
		meta, _ := window.DefaultWindow().CreateElement("meta", "")
		meta.SetAttribute("name", v)

		head := window.DefaultWindow().Get("document").
			Get("head")

		el, err := html.ToElement(head)
		if err != nil {
		}
		el.AppendChild(meta)
		return meta
	}

	ele, _ := html.ToElement(metaValue)
	return ele
}

func (p Page) metaByProperty(v string) js.Value {
	meta := window.DefaultWindow().
		Get("document").
		Call("querySelector", "meta[property='"+v+"']")

	if meta.IsNull() {
		metaEle, _ := window.DefaultWindow().CreateElement("meta", "")
		metaEle.SetAttribute("property", v)

		if head, err := html.ToElement(window.DefaultWindow().Get("document").
			Get("head")); err == nil {
			head.AppendChild(metaEle)
		}

	}

	return meta
}

func (p Page) Description() string {
	v, _ := p.metaByName("description").GetAttribute("content")
	return v
}

func (p Page) SetDescription(format string, v ...any) {
	description := fmt.Sprintf(format, v...)
	p.metaByName("description").SetAttribute("content", description)

	if ele, err := html.ToElement(p.metaByProperty("og:description")); err == nil {
		ele.SetAttribute("content", description)
	}

}

func (p Page) Author() string {
	v, _ := p.metaByName("author").GetAttribute("content")
	return v
}

func (p Page) SetAuthor(format string, v ...any) {
	p.metaByName("author").SetAttribute("content", fmt.Sprintf(format, v...))
}

func (p Page) Keywords() string {
	v, _ := p.metaByName("keywords").GetAttribute("content")
	return v
}

func (p Page) SetKeywords(v ...string) {
	p.metaByName("keywords").SetAttribute("content", strings.Join(v, ", "))
}

func (p Page) SetLoadingLabel(format string, v ...any) {
}

func (p Page) Preloads() []Preload {
	return nil
}

func (p Page) SetPreloads(v ...Preload) {
}

func (p Page) URL() *url.URL {
	return window.DefaultWindow().URL()
}
func (p Page) Size() (width int, height int) {
	return window.DefaultWindow().Size()
}
