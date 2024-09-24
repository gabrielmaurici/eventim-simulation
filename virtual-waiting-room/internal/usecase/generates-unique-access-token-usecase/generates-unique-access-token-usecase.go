package generates_unique_access_token_usecase

import (
	"github.com/gabrielmaurici/eventim-simulation/pkg/token"
)

type GeneratesUniqueAccessTokenOutputDTO struct {
	Token string `json:"token"`
}

type GeneratesUniqueAccessTokenUseCase struct {
}

func NewGeneratesUniqueAccessTokenUseCase() *GeneratesUniqueAccessTokenUseCase {
	return &GeneratesUniqueAccessTokenUseCase{}
}

func (uc *GeneratesUniqueAccessTokenUseCase) Execute() (*GeneratesUniqueAccessTokenOutputDTO, error) {
	token, err := token.GenerateAccessToken()

	if err != nil {
		return nil, err
	}

	return &GeneratesUniqueAccessTokenOutputDTO{
		Token: token,
	}, nil
}
