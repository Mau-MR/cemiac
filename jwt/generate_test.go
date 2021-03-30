package jwt

import (
	"github.com/Mau-MR/cemiac/data/auth"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestJWTManager_Generate(t *testing.T) {
	usr := auth.User{Name: "Mauricio", Surname: "Merida", Mail: "maupdtvty@gmail.com", Password: "somepass", Role: "admin"}
	manager := NewJWTManager("secretSouce", 15*time.Minute)
	tkn,err:= manager.Generate(&usr)
	assert.Nil(t, err,"ERROR ON THE  CREATION  OF THE  TOKEN")
	claims,err:= manager.Verify(tkn)
	assert.Nil(t, err,"ERROR ON THE VERIFICATION OF THE TOKEN")
	assert.Equal(t, &usr,claims.User,usr)
}

