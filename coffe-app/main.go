package main

import (
	"context"
	_ "fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"
	"github.com/shafey01/Go-Apps-Books/go-web-programming-Book/coffe-app/handlers"
)

func main() {
	l := log.New(os.Stdout, "product-api", log.LstdFlags)

	ph := handlers.NewProducts(l)

	// Multiplexer
	serverMux := mux.NewRouter()

	getRouter := serverMux.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc("/", ph.GetProducts)

	postRouter := serverMux.Methods(http.MethodPost).Subrouter()
	postRouter.HandleFunc("/", ph.AddProduct)
	postRouter.Use(ph.ValidateMiddelwareProducts)

	putRouter := serverMux.Methods(http.MethodPut).Subrouter()
	putRouter.HandleFunc("/{id:[0-9]+}", ph.UpdateProduct)
	putRouter.Use(ph.ValidateMiddelwareProducts)

	deleteRouter := serverMux.Methods(http.MethodDelete).Subrouter()
	deleteRouter.HandleFunc("/product/{id:[0-9]+}", ph.DeleteProduct)

	// server struct
	server := &http.Server{
		Addr:         "0.0.0.0:8080",
		Handler:      serverMux,
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
