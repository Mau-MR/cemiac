package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewSHA256(t *testing.T) {
	s := NewSHA256([]byte("hellothere"))
	assert.NotNil(t, s)
}
func TestEncryptString(t *testing.T) {
	s := EncryptString("hellothere")
	assert.NotNil(t, s)
}
