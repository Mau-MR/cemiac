package jwt

import (
	"github.com/Mau-MR/cemiac/data/auth"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestJWTManager_Verify(t *testing.T) {
	acc := NewJWTManager("secretsouce", 20*time.Minute)
	usr := auth.User{Mail: "Maudevops@gmail.com", Role: "admin", Name: "Mauricio Eulalio", Surname: "Merida Rivera", Password: "some pass"}
	tkn, err := acc.Generate(&usr)
	assert.Nil(t, err,"ERROR IN THE CREATION OF THE TOKEN")
	claims,err := acc.Verify(tkn)
	assert.Nil(t, err, "ERROR IN THE VERIFICATION OF THE TOKEN")
	assert.Equal(t,claims.User,&usr,"MODIFICATION OF THE ACCOUNT")

	//invalid jwt
	//TODO: ADD SOME VALIDATION FOR JWT AND THEIR SECURITY VULNERABILITIES
	claims,err = acc.Verify("someinvalid.token.gg")
	assert.NotNil(t, err)

}
