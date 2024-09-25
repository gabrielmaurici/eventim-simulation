package gateway

type BuyersActivesGateway interface {
	GetBuyersActives() (total int, err error)
	Add(token string) (err error)
}
