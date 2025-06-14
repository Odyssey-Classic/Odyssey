package registry

import (
	"context"
	"log/slog"
	"time"
)

func (r *Registry) Start(ctx context.Context) error {
	if err := r.Ping(ctx); err != nil {
		return err
	}

	r.once.Do(func() {
		r.wg.Add(1)
		go func() {
			r.start(ctx)
			r.wg.Done()
		}()
	})

	return nil
}

func (r *Registry) start(ctx context.Context) {
	slog.Info("[registry] service starting")
	var err error
	for {
		select {
		case <-ctx.Done():
			slog.Info("[regsitry] stopping")
			return
		default:
			time.Sleep(5 * time.Second)
			err = r.Ping(ctx)
			if err != nil {
				slog.Error("[registry] error pinging", "error", err)
			}
		}
	}
}
