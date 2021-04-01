package jwt

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestNewJWTManager(t *testing.T) {
	manager:= NewJWTManager("secrretasouce",20*time.Minute)
	assert.NotNil(t, manager,"THE INITIALIZATION OF THE JWT MANAGER WAS NIL")
}
