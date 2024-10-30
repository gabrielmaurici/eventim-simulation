package mocks

import (
	"context"

	"github.com/stretchr/testify/mock"
)

type BuyersActivesGatewayMock struct {
	mock.Mock
}

func (m *BuyersActivesGatewayMock) GetBuyersActives(ctx context.Context) (total int64, err error) {
	args := m.Called(ctx)
	return int64(args.Int(0)), args.Error(1)
}

func (m *BuyersActivesGatewayMock) Add(token string, ctx context.Context) error {
	args := m.Called(token, ctx)
	return args.Error(0)
}

func (m *BuyersActivesGatewayMock) Delete(userToken string, ctx context.Context) error {
	args := m.Called(userToken, ctx)
	return args.Error(0)
}
