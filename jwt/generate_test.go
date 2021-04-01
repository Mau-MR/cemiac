package jwt

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestJWTManager_Generate(t *testing.T) {
	acc := &UserInfo{
		Name: "Mauricio",
		Surname: "Merida",
		Role: "admin",
		ID: "2442JALDJGABNLAJ",

	}
	manager := NewJWTManager("secretSouce", 15*time.Minute)
	tkn,err:= manager.Generate(acc)
	assert.Nil(t, err,"ERROR ON THE  CREATION  OF THE  TOKEN")
	claims,err:= manager.Verify(tkn)
	assert.Nil(t, err,"ERROR ON THE VERIFICATION OF THE TOKEN")
	assert.Equal(t, acc,claims.UserInfo,"They are not  equal" )
}

