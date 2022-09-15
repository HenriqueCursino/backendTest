package common

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConstants(t *testing.T) {
	assert.Equal(t, STATUS_PENDENTE, 1)
	assert.Equal(t, STATUS_CONCLUIDO, 2)
	assert.Equal(t, LOJISTA, 1)
}
