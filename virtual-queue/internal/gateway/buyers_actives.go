package gateway

type BuyersActivesGateway interface {
	GetBuyersActives() (total int64, err error)
	Add(token string) (err error)
}
