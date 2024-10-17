package database

import (
	"database/sql"
	"fmt"

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

func (t *TicketDb) GetAvailableTickets(quantity int8) ([]entity.Ticket, error) {
	rows, err := t.Db.Query("SELECT id, available FROM tickets WHERE available = TRUE LIMIT ?", quantity)
	if err != nil {
		return nil, fmt.Errorf("erro ao consultar ingressos: %v", err)
	}
	defer rows.Close()

	var tickets []entity.Ticket
	for rows.Next() {
		var ticket entity.Ticket
		if err := rows.Scan(&ticket.Id, &ticket.Available); err != nil {
			return nil, fmt.Errorf("erro ao ler dados dos ingressos: %v", err)
		}
		tickets = append(tickets, ticket)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("erro durante iteração dos ingressos: %v", err)
	}

	return tickets, nil
}

func (t *TicketDb) Update(ticket entity.Ticket) error {
	stmt, err := t.Db.Prepare("UPDATE tickets SET available = ? WHERE id = ?")
	if err != nil {
		return fmt.Errorf("erro ao preparar query de update: %v", err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(ticket.Available, ticket.Id)
	if err != nil {
		return fmt.Errorf("erro ao executar update do ingresso: %v", err)
	}

	return nil
}
