package firestore

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateClient(t *testing.T) {
	client :=CreateClient()
	assert.NotNil(t, client,"The  client  should not be nil")
}
