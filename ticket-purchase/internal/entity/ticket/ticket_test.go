package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_WhenInputIsValid_ReturnNewTicket(t *testing.T) {
	id := "123"
	available := true

	ticket := NewTicket(id, available)

	assert.Equal(t, id, ticket.Id)
	assert.Equal(t, available, ticket.Available)
}

func Test_WhenUpdateToUnavailable_FieldAvaiableReturnsFalse(t *testing.T) {
	id := "123"
	available := true

	ticket := NewTicket(id, available)
	ticket.UpdateToUnavailable()

	assert.Equal(t, id, ticket.Id)
	assert.False(t, ticket.Available)
}
