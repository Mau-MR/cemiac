package main

import (
	"context"
	"github.com/Mau-MR/cemiac/firestore"
	"github.com/Mau-MR/cemiac/handlers"
	"github.com/Mau-MR/cemiac/utils"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	l := log.New(os.Stdout, "cemiac-api ", log.LstdFlags)
	dbClient := firestore.CreateClient()
	validation := utils.NewValidation()

	users := handlers.NewUsers(l,validation,dbClient)


	mux := mux.NewRouter()
	//Post methods
	postRouter := mux.Methods(http.MethodPost).Subrouter()
	postRouter.HandleFunc("/login",users.Login)

	server := http.Server{
		Addr:         "localhost:8080",
		Handler:      mux,
		ErrorLog:     l,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  100 * time.Second,
	}



	go func() {
		l.Println("Starting server on port 8080")
		if err := server.ListenAndServe(); err !=nil{
			l.Fatal("Error starting the sever: ", err)
			os.Exit(1)
		}

	}()
	// trap sigterm or interupt and gracefully shutdown the server

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, os.Kill)

	// Block until a signal is received.
	sig := <-c
	log.Println("Got signal:", sig)

	// gracefully shutdown the server, waiting max 30 seconds for current operations to complete
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	if err := server.Shutdown(ctx); err!=nil{
		l.Fatal("Error turning down the server: ", err)
	}

}
