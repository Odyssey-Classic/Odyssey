package network

import (
	"context"
	"log/slog"

	"golang.org/x/sync/errgroup"
)

func (n *Network) processClient(ctx context.Context, client *Client) {
	eg, ctx := errgroup.WithContext(ctx)

	n.clientGroup.Add(1)
	go func() {
		// errgroup.Go does not have a way to use its own context
		eg.Go(func() error { return client.processInbound(ctx) })
		eg.Go(func() error { return client.processOutbound(ctx) })

		err := eg.Wait()
		slog.Error("client process failed", "error", err)
		n.clientGroup.Done()
	}()
}
