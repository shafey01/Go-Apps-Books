package handlers

import (
	"log"
	"net/http"

	"github.com/shafey01/Go-Apps-Books/go-web-programming-Book/chitchat/data"
)

type Products struct {
	l *log.Logger
}

func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

// Method serverHttp to staisfies the http.Handler
// New Version
func (p *Products) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	// handle Get request
	if r.Method == http.MethodGet {
		p.getProducts(w, r)
		return
	}
	if r.Method == http.MethodDelete {
		p.deleteProduct(w, r)
		return
	}
	if r.Method == http.MethodPost {
		p.addProduct(w, r)
		return
	}
	// else return an error
	w.WriteHeader(http.StatusMethodNotAllowed)

}

func (p *Products) getProducts(w http.ResponseWriter, r *http.Request) {

	p.l.Println("Handle Get Product Request")

	// fetch data from the database
	listProduct := data.GetProducts()
	err := listProduct.ToJSON(w)
	if err != nil {
		http.Error(w, "Unable to marshal json", http.StatusInternalServerError)
	}

}

func (p *Products) addProduct(w http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle Post Request")
	NewProduct := &data.Product{}
	err := NewProduct.FromJSON(r.Body)
	if err != nil {
		http.Error(w, "Unable to Decode data", http.StatusBadRequest)

	}

	data.AddProduct(NewProduct)
}

func (p *Products) deleteProduct(w http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle delete Request")
}
