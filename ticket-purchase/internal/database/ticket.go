package database

import (
	"database/sql"

	entity "github.com/gabrielmaurici/eventim-simulation/ticket-purchase/internal/entity/ticket"
)

type TicketDb struct {
	Db *sql.DB
}

func NewTicketDb(db *sql.DB) *TicketDb {
	return &TicketDb{
		Db: db,
	}
}

func (t *TicketDb) Get(id string) (*entity.Ticket, error) {
	var ticket entity.Ticket
	err := t.Db.QueryRow("SELECT id, available FROM tickets WHERE id = ?", id).Scan(&ticket.Id, &ticket.Available)
	if err != nil {
		return nil, err
	}

	return &ticket, nil
}

func (t *TicketDb) GetAvailableTickets(quantity int8) (*[]entity.Ticket, error) {
	rows, err := t.Db.Query("SELECT id, available FROM tickets WHERE available = TRUE LIMIT ?", quantity)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tickets []entity.Ticket
	for rows.Next() {
		var ticket entity.Ticket
		if err := rows.Scan(&ticket.Id, &ticket.Available); err != nil {
			return nil, err
		}
		tickets = append(tickets, ticket)
	}

	return &tickets, nil
}

func (t *TicketDb) Update(ticket *entity.Ticket) error {
	stmt, err := t.Db.Prepare("UPDATE tickets SET available = ? WHERE id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(ticket.Available, ticket.Id)
	if err != nil {
		return err
	}

	return nil
}
