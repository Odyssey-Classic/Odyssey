package admin

import (
	"context"
	"log/slog"
	"sync"
)

type Admin struct {
	wg *sync.WaitGroup
}

func New(wg *sync.WaitGroup) *Admin {
	return &Admin{wg: wg}
}

func (a *Admin) Start(ctx context.Context) {
	a.wg.Add(1)
	go func() {
		a.start(ctx)
		a.wg.Done()
	}()
}

func (a *Admin) start(ctx context.Context) {
	for {
		<-ctx.Done()
		slog.Info("admin shutting down")
		return
	}
}
