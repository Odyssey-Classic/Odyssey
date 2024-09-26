package game

import (
	"context"
	"log/slog"
	"sync"
)

type Game struct {
	wg *sync.WaitGroup
}

func New(wg *sync.WaitGroup) *Game {
	return &Game{wg: wg}
}

func (g *Game) Start(ctx context.Context, network chan any) {
	g.wg.Add(1)
	go func() {
		g.start(ctx, network)
		g.wg.Done()
	}()
}

func (g *Game) start(ctx context.Context, network chan any) {
	for {
		select {
		case msg := <-network:
			slog.Info("game received message: ", "message", msg)
			// msg.Type = new Client
			//   game.NewPlayer(client)
		case <-ctx.Done():
			slog.Info("game shutting down")
			return
		}
	}
}
