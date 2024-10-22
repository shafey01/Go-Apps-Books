package handlers

import (
	"log"
	"net/http"
	"github.com/shafey01/Go-Apps-Books/tree/main/go-web-programming-Book/chitchat/data"
)

type Products struct {
	l *log.Logger
}

func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

// Method serverHttp to staisfies the http.Handler
func (p *Products) ServeHttp(w http.ResponseWriter, r *http.Request) {

	// handle Get request
	if r.Method == http.MethodGet {
		p.getProducts(w, r)
		return
	}

	// else return an error
	w.WriteHeader(http.StatusMethodNotAllowed)

}

func (p *Products) getProducts(w http.ResponseWriter, r *http.Request) {

	p.l.Println("Handle Get Product Request")

	// fetch data from the database
	listProduct := data.GetProducts()





}
