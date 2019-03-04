// Package micro provides a micro rpc resolver which prefixes a namespace
package micro

import (
	"net/http"

	"github.com/micro/go-api/resolver"
)

// default resolver for legacy purposes
// it uses proxy routing to resolve names
// /foo becomes namespace.foo
// /v1/foo becomes namespace.v1.foo
type microResolver struct {
	opts resolver.Options
}

func (r *microResolver) Resolve(req *http.Request) (*resolver.Endpoint, error) {
	var name string
	var method string

	switch r.opts.Handler {
	case "meta", "api", "rpc":
		name, method = apiRoute(r.opts.Namespace, req.URL.Path)
	default:
		method = req.Method
		name = proxyRoute(r.opts.Namespace, req.URL.Path)
	}

	return &resolver.Endpoint{
		Name:   name,
		Method: method,
	}, nil
}

func (r *microResolver) String() string {
	return "micro"
}

// NewResolver creates a new micro resolver
func NewResolver(opts ...resolver.Option) resolver.Resolver {
	return &microResolver{
		opts: resolver.NewOptions(opts...),
	}
}
