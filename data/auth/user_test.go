package auth

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewUser(t *testing.T) {
	user := NewUser("mauricio","Merida","maudevops@gmail.com","somepass","admin")
	assert.NotNil(t, user)
}
