package vertex

import (
	"context"
	"net/url"
)

type Context struct {
	context.Context
	navigate func(*url.URL, bool)
	dispatch func(func())
	defere   func(func())
}

// Navigate transitions to the given URL string.
func (ctx Context) Navigate(rawURL string) {
	u, err := url.Parse(rawURL)
	if err != nil {
		//Log(errors.New("navigating to URL failed").
		//	WithTag("url", rawURL).
		//	Wrap(err))
		return
	}
	ctx.NavigateTo(u)
}

// NavigateTo transitions to the provided URL.
func (ctx Context) NavigateTo(u *url.URL) {
	ctx.navigate(u, true)
}
