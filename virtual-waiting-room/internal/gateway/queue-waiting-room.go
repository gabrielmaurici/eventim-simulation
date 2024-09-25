package gateway

type QueueWaitingRoomGateway interface {
	Enqueue(token string) (position int, err error)
}
