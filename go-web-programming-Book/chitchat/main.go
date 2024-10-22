package main

import (
	"context"
	_ "fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/shafey01/Go-Apps-Books/go-web-programming-Book/chitchat/handlers"
)

func main() {

	l := log.New(os.Stdout, "product-api", log.LstdFlags)

	// Multiplexer
	mux := http.NewServeMux()

	ph := handlers.NewProducts(l)

	// mux handler function
	mux.Handle("/", ph)

	// server struct
	server := &http.Server{
		Addr:         "0.0.0.0:8080",
		Handler:      mux,
		ErrorLog:     l,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	// go routine to start the server

	go func() {
		l.Println("Starting the server on port 8080")

		err := server.ListenAndServe()
		if err != nil {
			l.Printf("Erorr starting the server %s \n", err)
			os.Exit(1)
		}
	}()
	// send signals to channel to exit softly
	errorChannel := make(chan os.Signal, 1)
	signal.Notify(errorChannel, os.Interrupt)
	signal.Notify(errorChannel, os.Kill)

	// block by channel until signal received
	sig := <-errorChannel
	log.Println("Got signal: ", sig)

	// shutdown gracefully
	ctx, _ := context.WithTimeout(context.Background(), 3*time.Second)
	server.Shutdown(ctx)

}
