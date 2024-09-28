package token

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateAccessToken(t *testing.T) {
	expectedLength := 64

	token, _ := GenerateUniqueAccessToken()

	assert.Equal(t, expectedLength, len(token))
}

func TestGenerateMultipleTokens(t *testing.T) {
	token1, _ := GenerateUniqueAccessToken()
	token2, _ := GenerateUniqueAccessToken()

	assert.NotEqual(t, token1, token2)
}
