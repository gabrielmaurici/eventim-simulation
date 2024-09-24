package token

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateAccessToken(t *testing.T) {
	expectedLength := 64

	token, _ := GenerateAccessToken()

	assert.Equal(t, expectedLength, len(token))
}

func TestGenerateMultipleTokens(t *testing.T) {
	token1, _ := GenerateAccessToken()
	token2, _ := GenerateAccessToken()

	assert.NotEqual(t, token1, token2)
}
