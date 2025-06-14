package registry

import (
	"context"
	"encoding/base64"
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

func New(wg *sync.WaitGroup, url *url.URL, id string, apiKey string) *Registry {
	auth := base64.StdEncoding.EncodeToString([]byte(id + ":" + apiKey))
	rt := &HeaderRoundTripper{
		Headers: map[string]string{
			"Authorization": "Basic " + auth,
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
