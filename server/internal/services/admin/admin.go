package admin

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"sync"
	"time"

	"github.com/go-chi/chi/v5"
)

type Admin struct {
	wg   *sync.WaitGroup
	port uint16
	once sync.Once
}

func New(wg *sync.WaitGroup, port uint16) *Admin {
	return &Admin{wg: wg, port: port}
}

func (a *Admin) Start(ctx context.Context) {
	a.once.Do(func() {
		a.wg.Add(1)
		go func() {
			a.start(ctx)
			a.wg.Done()
		}()
	})
}

func (a *Admin) start(ctx context.Context) {
	r := chi.NewRouter()

	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("ok"))
	})

	srv := &http.Server{
		Addr:    ":" + fmt.Sprintf("%d", a.port), // Use the port from the Admin struct
		Handler: r,
	}

	go func() {
		<-ctx.Done()
		slog.Info("admin shutting down")
		shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		if err := srv.Shutdown(shutdownCtx); err != nil {
			slog.Error("admin shutdown error", "err", err)
		} else {
			slog.Info("admin shutdown complete")
		}
	}()

	slog.Info("admin API starting on :" + fmt.Sprintf("%d", a.port))
	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		slog.Error("admin API server error", "err", err)
	}
}
