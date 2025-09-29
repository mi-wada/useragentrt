// Package useragentrt provides a http.RoundTripper that sets a User-Agent header.
package useragentrt

import "net/http"

var _ http.RoundTripper = (*roundTripper)(nil)

type roundTripper struct {
	base http.RoundTripper
	ua   string
}

// New returns a [http.RoundTripper] that sets ua to the User-Agent header.
func New(base http.RoundTripper, ua string) *roundTripper {
	if base == nil {
		base = http.DefaultTransport
	}
	return &roundTripper{base: base, ua: ua}
}

// RoundTrip implements the [http.RoundTripper] interface.
func (rt *roundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	// Clone the request to avoid modifying the original request.
	// https://cs.opensource.google/go/go/+/refs/tags/go1.23.6:src/net/http/client.go;l=128-132
	clonedReq := req.Clone(req.Context())
	clonedReq.Header.Set("User-Agent", rt.ua)
	return rt.base.RoundTrip(clonedReq)
}
