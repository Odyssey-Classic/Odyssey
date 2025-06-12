package registry

import (
	"context"
	"net/http"
	"net/url"
	"sync"
)

type Registry struct {
	url *url.URL

	once sync.Once
	wg   *sync.WaitGroup
}

func New(wg *sync.WaitGroup, url *url.URL) *Registry {
	return &Registry{
		wg:  wg,
		url: url,
	}
}

func (r *Registry) Ping(ctx context.Context) error {
	_, err := http.Get(r.url.String())
	return err
}
