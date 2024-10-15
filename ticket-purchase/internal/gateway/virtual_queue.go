package gateway

type VirtualQueueGateway interface {
	Enqueue(token string) (position int64, err error)
	Dequeue() (token string, err error)
	GetAll() (tokens []string, err error)
}
