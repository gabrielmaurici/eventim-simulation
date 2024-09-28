package gateway

type VirtualQueueGateway interface {
	Enqueue(token string) (position int64, err error)
}
