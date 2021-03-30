package handlers

import (
	"cloud.google.com/go/firestore"
	"github.com/Mau-MR/cemiac/data/auth"
	"github.com/Mau-MR/cemiac/utils"
	"log"
	"net/http"
)

type Users struct{
	logger *log.Logger
	validation *utils.Validation
	usersDB *auth.AuthDB
}

func NewUsers(log *log.Logger, val *utils.Validation, client *firestore.Client)*Users{
	return &Users{
		logger: log,
		validation: val,
		usersDB: auth.NewAuthDB(log,client),
	}
}

func (u *Users) Login(rw http.ResponseWriter, r *http.Request)  {
	account := &auth.Account{}
	//Parse the  request to json and return the err in case of decoding error
	err := utils.ParseRequest(account,r.Body,rw)
	if err !=nil{
		u.logger.Println("Error parsing account:  ",err)
		return
	}
	//Checking the validation for  all the required fields
	err = u.validation.ValidateRequest(account,rw)
	if  err!=nil{
		u.logger.Println("Invalid arguments on the account" , account)
		return
	}
	//make the db call to verify  existance of  account

}
