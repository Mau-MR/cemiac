package auth

import (
	"cloud.google.com/go/firestore"
	"log"
)

type AuthDB struct {
	logger *log.Logger
	client *firestore.Client
}

func NewAuthDB(l *log.Logger,c *firestore.Client)  *AuthDB{
	return &AuthDB{
		logger: l,
		client: c,
	}
}

