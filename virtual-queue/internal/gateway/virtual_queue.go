package gateway

import "context"

type VirtualQueueGateway interface {
	Enqueue(token string, ctx context.Context) (position int64, err error)
	Dequeue(ctx context.Context) (token string, err error)
	GetAll(ctx context.Context) (tokens []string, err error)
}
