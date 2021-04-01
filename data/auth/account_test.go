package auth

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func  TestNewAccount(t *testing.T) {
	acc := NewAccount("maudevops@gmail.com","somepass")
	assert.NotNil(t, acc)
}