package mocks

import (
	"github.com/stretchr/testify/mock"
)

type ProducerMock struct {
	mock.Mock
}

func (p *ProducerMock) Publish(msg interface{}) error {
	args := p.Called(msg)
	return args.Error(0)
}
