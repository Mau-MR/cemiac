package handlers

import (
	"cloud.google.com/go/firestore"
	"github.com/Mau-MR/cemiac/data/auth"
	"github.com/Mau-MR/cemiac/jwt"
	"github.com/Mau-MR/cemiac/utils"
	"log"
	"net/http"
)

type Users struct {
	logger     *log.Logger
	validation *utils.Validation
	usersDB    *auth.AuthDB
	jwtManager *jwt.JWTManager
}

func NewUsers(log *log.Logger, val *utils.Validation, client *firestore.Client, manager *jwt.JWTManager) *Users {
	return &Users{
		logger:     log,
		validation: val,
		usersDB:    auth.NewAuthDB(log, client),
		jwtManager: manager,
	}
}

//TODO: CHANGE THE ERROR RETURN TO SOME GENERIC ERROR TO HIDE INFORMATION OF THE INTERNAL INFRASTRUCTURE
//Returns the token in case of correct  account
func (u *Users) Login(rw http.ResponseWriter, r *http.Request) {
	account := &auth.Account{}
	//Parse the  request to json and return the err in case of decoding error
	err := utils.ParseRequest(account, r.Body, rw)
	if err != nil {
		u.logger.Println("Error parsing account:  ", err)
		return
	}
	//Checking the validation for  all the required fields
	err = u.validation.ValidateRequest(account, rw)
	if err != nil {
		u.logger.Println("Invalid arguments on the account", account)
		return
	}
	//TODO REFACTOR THIS  CODE  TO  MAKE  IT  CLEANER
	user, err := u.usersDB.Login(account)
	if err != nil {
		u.logger.Println("Error searching on the db", err)
		rw.WriteHeader(http.StatusNotFound)
		utils.ToJSON(utils.GenericError{
			Message: err.Error(),
		}, rw)
		return
	}
	tkn, err := u.jwtManager.Generate(user)
	if err != nil {
		u.logger.Println("Unable to generate the token for that account")
		rw.WriteHeader(http.StatusExpectationFailed)
		utils.ToJSON(utils.GenericError{
			Message: err.Error(),
		}, rw)
		return
	}
	rw.WriteHeader(http.StatusOK)
	//TODO:  CHECK FOR THE  VALIDATION  OF  THIS WITH THE REFACTOR OF THE ERROR STATUS
	utils.ToJSON(jwt.NewToken(tkn), rw)
	log.Println(tkn)
	return
}
func (u *Users) CreateAccount(rw http.ResponseWriter, r *http.Request) {

}
