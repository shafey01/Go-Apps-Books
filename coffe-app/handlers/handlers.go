package handlers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/shafey01/Go-Apps-Books/go-web-programming-Book/coffe-app/data"
)

type Products struct {
	l *log.Logger
}

func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

// Method serverHttp to staisfies the http.Handler
// New Version

func (p *Products) GetProducts(w http.ResponseWriter, r *http.Request) {

	p.l.Println("Handle Get Product Request")

	// fetch data from the database
	listProduct := data.GetProducts()
	err := listProduct.ToJSON(w)
	if err != nil {
		http.Error(w, "Unable to marshal json", http.StatusInternalServerError)
	}

}

func (p *Products) AddProduct(w http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle Post Request")
	NewProduct := &data.Product{}
	err := NewProduct.FromJSON(r.Body)
	if err != nil {
		http.Error(w, "Unable to Decode data", http.StatusBadRequest)
		return
	}

	data.AddProduct(NewProduct)
}

func (p *Products) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {

		http.Error(w, "Unable to convert id", http.StatusBadRequest)
		return
	}
	p.l.Println("Handle PUT Request")
	NewProduct := &data.Product{}
	err = NewProduct.FromJSON(r.Body)
	if err != nil {
		http.Error(w, "Unable to Decode data", http.StatusBadRequest)

	}

	err = data.UpdateProduct(id, NewProduct)
	if err != nil {
		http.Error(w, "Product Not Found", http.StatusInternalServerError)
		return
	}
	if err == data.ErrProductNotFound {

		http.Error(w, "Product Not Found", http.StatusNotFound)
		return
	}
}
func (p *Products) deleteProduct(w http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle delete Request")
}
