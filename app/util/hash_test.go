package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const randomString = "somerandompassword"

func TestHashPassword(t *testing.T) {
	hash, err := HashPassword(randomString)
	if err != nil {
		assert.Fail(t, err.Error())
	}
	assert.True(t, CheckPasswordHash("somerandompassword", hash))
}
