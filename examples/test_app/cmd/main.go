package main

import (
	"github.com/volts-dev/volts"
	"github.com/volts-dev/volts/router"
	"github.com/volts-dev/volts/server"
	"github.com/volts-dev/volts/transport"
)

func main() {
	r := router.New()
	r.Url("GET", "/", func(c *router.THttpContext) {
		c.RenderTemplate("index.html", nil)
		//c.ServeFile("index.html")
	})

	tr := transport.NewHTTPTransport(
		transport.Addrs(":9999"),
		transport.Debug(),
	)

	srv := volts.New()
	srv.Server().Config().Init(
		server.WithTransport(tr),
		server.WithRouter(r),
	)

	srv.Run()
}
