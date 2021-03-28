package main

import (
	"context"
	"github.com/Mau-MR/cemiac/handlers"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)
func HelloHandler(rw http.ResponseWriter, r *http.Request){
	rw.Write([]byte("Hello"))
}
func main() {
	l := log.New(os.Stdout, "cemiac-api", log.LstdFlags)
	mux := http.NewServeMux()
	mux.HandleFunc("/v1/hello",HelloHandler)

	wrappedMux := handlers.NewAuthMiddleware(mux,l)
	server := http.Server{
		Addr:         "localhost:8080",
		Handler:      wrappedMux,
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
