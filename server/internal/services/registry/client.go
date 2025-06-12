package registry

import "net/http"

// HeaderRoundTripper adds default headers to every request.
type HeaderRoundTripper struct {
	Headers   map[string]string
	Transport http.RoundTripper
}

func (h *HeaderRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	for k, v := range h.Headers {
		req.Header.Set(k, v)
	}
	return h.transport().RoundTrip(req)
}

func (h *HeaderRoundTripper) transport() http.RoundTripper {
	if h.Transport != nil {
		return h.Transport
	}
	return http.DefaultTransport
}
