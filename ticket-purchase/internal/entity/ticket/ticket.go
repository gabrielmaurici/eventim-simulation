package entity

type Ticket struct {
	Id        string `json:"id"`
	Available bool   `json:"available"`
}

func NewTicket(id string, available bool) *Ticket {
	return &Ticket{
		Id:        id,
		Available: available,
	}
}

func (t *Ticket) UpdateToUnavailable() {
	t.Available = false
}
