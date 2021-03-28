package handlers

import (
	"log"
	"net/http"
)

type Auth struct {
	Handler http.Handler
	Logger  log.Logger
}

func NewAuthMiddleware(handlerToWrap http.Handler, l *log.Logger) *Auth {
	return &Auth{
		Handler: handlerToWrap,
		Logger: *l,
	}
}

func (a *Auth) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	log.Printf("%s %s ", r.Method, r.URL.Path)
	a.Handler.ServeHTTP(rw,r)
}
