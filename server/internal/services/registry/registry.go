package registry

import (
	"context"
	"net/http"
	"net/url"
	"sync"
)

type Registry struct {
	url    *url.URL
	client *http.Client

	once sync.Once
	wg   *sync.WaitGroup
}

func New(wg *sync.WaitGroup, url *url.URL, apiKey string) *Registry {
	rt := &HeaderRoundTripper{
		Headers: map[string]string{
			"Authorization": apiKey,
		},
	}
	client := &http.Client{Transport: rt}
	return &Registry{
		wg:     wg,
		url:    url,
		client: client,
	}
}

func (r *Registry) Ping(ctx context.Context) error {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, r.url.String()+"/server", nil)
	if err != nil {
		return err
	}
	_, err = r.client.Do(req)
	return err
}
