package gateway

import "context"

type BuyersActivesGateway interface {
	GetBuyersActives(ctx context.Context) (total int64, err error)
	Add(token string, ctx context.Context) error
	Delete(token string, ctx context.Context) error
}
