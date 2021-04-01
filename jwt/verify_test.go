package jwt

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestJWTManager_Verify(t *testing.T) {
	acc := NewJWTManager("secretsouce", 20*time.Minute)
	info := &UserInfo{Name: "mauricio",Surname: "Merida",Role: "admin",ID: "akajqkl2351"}
	tkn, err := acc.Generate(info)
	assert.Nil(t, err,"ERROR IN THE CREATION OF THE TOKEN")
	claims,err := acc.Verify(tkn)
	assert.Nil(t, err, "ERROR IN THE VERIFICATION OF THE TOKEN")
	assert.Equal(t,claims.UserInfo,info,"MODIFICATION OF THE ACCOUNT")

	//invalid jwt
	//TODO: ADD SOME VALIDATION FOR JWT AND THEIR SECURITY VULNERABILITIES
	claims,err = acc.Verify("someinvalid.token.gg")
	assert.NotNil(t, err)

}
