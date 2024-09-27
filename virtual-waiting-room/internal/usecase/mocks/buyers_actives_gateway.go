package mocks

import "github.com/stretchr/testify/mock"

type BuyersActivesGatewayMock struct {
	mock.Mock
}

func (m *BuyersActivesGatewayMock) GetBuyersActives() (total int64, err error) {
	args := m.Called()
	return int64(args.Int(0)), args.Error(1)
}

func (m *BuyersActivesGatewayMock) Add(token string) (err error) {
	args := m.Called(token)
	return args.Error(0)
}
