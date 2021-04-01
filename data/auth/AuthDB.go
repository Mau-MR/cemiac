
package auth

import (
	"cloud.google.com/go/firestore"
	"context"
	"fmt"
	"github.com/Mau-MR/cemiac/jwt"
	"github.com/Mau-MR/cemiac/utils"
	"log"
	"time"
)

type AuthDB struct {
	logger     *log.Logger
	client     *firestore.Client
	collection string
}

func NewAuthDB(l *log.Logger, c *firestore.Client) *AuthDB {
	return &AuthDB{
		logger:     l,
		client:     c,
		collection: "Users",
	}
}

//TODO check if its necesary to add the firestore tags to the struct and test the insertion of this
func (a *AuthDB) CreateUser(user *User) error {
	usr, err := a.AccountByMail(user.Mail)
	if err != nil {
		return err
	}
	//preventing the  account repetition
	if usr!= nil {
		return fmt.Errorf("User already exist")
	}
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	//Password encryption
	user.Password = utils.EncryptString(user.Password)
	ref := a.client.Collection(a.collection).NewDoc()
	_,err = ref.Set(ctx,user)
	if err != nil {
		a.logger.Println("Error inserting to db the new account ")
		return err
	}
	return nil
}

func (a *AuthDB) Login(account *Account) (*jwt.UserInfo, error) {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	//TODO:
	docs, err := a.client.Collection(a.collection).Where("mail", "==", account.Mail).Limit(1).Documents(ctx).GetAll()
	if err != nil {
		return nil, err
	}
	var userinfo jwt.UserInfo
	for _, doc := range docs {
		docMap := doc.Data()
		if docMap["password"]!=utils.EncryptString(account.Password){
			return nil,fmt.Errorf("incorrect password")
		}
		if err := doc.DataTo(&userinfo); err != nil {
			a.logger.Println(userinfo)
			return nil, err
		}
		userinfo.ID= doc.Ref.ID
		return &userinfo,nil
	}
	//Means there was no document with that mail
	return  nil,fmt.Errorf("Incorrect Mail")
}

//Searchs for the user with specific email and  returns the information of the user
func (a *AuthDB) AccountByMail(mail string) (*User, error) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	//TODO:
	docs, err := a.client.Collection(a.collection).Where("mail", "==", mail).Limit(1).Documents(ctx).GetAll()
	var user User
	if err != nil {
		return nil, err
	}
	for _, doc := range docs {
		if err := doc.DataTo(&user); err != nil {
			return nil, err
		}
		return &user, nil
	}
	return nil,nil
}
