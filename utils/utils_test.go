package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIDStringToUint64_Success(t *testing.T) {
	result := IDStringToUint64("1")
	assert.Equal(t, uint64(1), result)
}

func TestIDStringToUint64_Error(t *testing.T) {
	result := IDStringToUint64("-1")
	assert.Equal(t, uint64(0), result)

	result = IDStringToUint64("a")
	assert.Equal(t, uint64(0), result)
}

func TestHashBCrypt(t *testing.T) {
	hashed, err := HashBCrypt("superA@999")
	assert.NoError(t, err)
	assert.NotEmpty(t, hashed)
}
