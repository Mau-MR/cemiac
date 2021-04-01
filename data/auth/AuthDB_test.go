package auth

import (
	"github.com/stretchr/testify/assert"
	"log"
	"os"
	"testing"
	"github.com/Mau-MR/cemiac/firestore"
)

func TestNewAuthDB(t *testing.T) {
	db := NewAuthDB(&log.Logger{},firestore.CreateClient())
	assert.NotNil(t, db)
}
func TestAuthDB_CreateUser(t *testing.T) {
	usr:= NewUser("Mauricio Eulalio","Merida Rivera","maudevops@gmail.com","AntaliaM88!","admin")
	db := NewAuthDB(log.New(os.Stdout,"test",log.LstdFlags),firestore.CreateClient())
	err := db.CreateUser(usr)
	log.Println(err)
	assert.Nil(t, err,"This user had already been registered to db ")

}
func TestAuthDB_AccountByMail(t *testing.T) {
	db := NewAuthDB(log.New(os.Stdout,"test",log.LstdFlags),firestore.CreateClient())
	usr,err := db.AccountByMail("test")
	assert.Nil(t, err,"Error was not nil")
	log.Println(usr)
}
func TestAuthDB_Login(t *testing.T) {
	db := NewAuthDB(log.New(os.Stdout,"test",log.LstdFlags),firestore.CreateClient())
	acc := &Account{Mail: "maudevops@gmail.com",Password: "AntaliaM88!"}
	user,err := db.Login(acc)
	assert.Nil(t, err)
	log.Println(user)
}
