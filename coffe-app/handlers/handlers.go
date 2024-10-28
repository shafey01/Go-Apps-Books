package handlers

import (
	"context"
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

	if len(listProduct) == 0 {

		p.l.Println("Handle Get Product Request: [ERROR] Empty List")
		http.Error(w, "Empty List", http.StatusBadRequest)
		return
	}
	err := listProduct.ToJSON(w)
	if err != nil {
		http.Error(w, "Unable to marshal json", http.StatusInternalServerError)
	}

}

func (p *Products) AddProduct(w http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle Post Request")
	NewProduct := r.Context().Value(KeyProduct{}).(data.Product)
	data.AddProduct(&NewProduct)
}

func (p *Products) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {

		http.Error(w, "Unable to convert id", http.StatusBadRequest)
		return
	}
	p.l.Println("Handle PUT Request")
	NewProduct := r.Context().Value(KeyProduct{}).(data.Product)
	err = data.UpdateProduct(id, &NewProduct)
	if err == data.ErrProductNotFound {
		p.l.Println("[ERORR] Product Not Found ")
		http.Error(w, "Product Not Found", http.StatusNotFound)
		w.WriteHeader(http.StatusNotFound)
		return
	}
}
func (p *Products) DeleteProduct(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {

		http.Error(w, "Unable to convert id", http.StatusBadRequest)
		return
	}

	p.l.Println("Handle delete Request")
	err = data.DelelteProduct(id)

	if err == data.ErrProductNotFound {

		p.l.Println("[ERORR] Product Not Found ")
		w.WriteHeader(http.StatusNotFound)
		return
	}

	if err == data.ErrEmptyList {

		p.l.Println("[ERORR] Product Not Found ")
		w.WriteHeader(http.StatusNotFound)
		return
	}
}

// Key type for context
type KeyProduct struct{}

// Middleware function to validate the request and call the next handler
func (p Products) ValidateMiddelwareProducts(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		NewProduct := data.Product{}
		err := NewProduct.FromJSON(r.Body)
		if err != nil {
			p.l.Println("[ERROR] deserializing product", err)
			http.Error(w, "Unable to Decode data", http.StatusBadRequest)
			return
		}
		err = NewProduct.Validate()
		if err != nil {

			p.l.Println("[ERROR] Validating product", err)
			http.Error(w, "Unable to Validating data", http.StatusBadRequest)
			return

		}
		ctx := context.WithValue(r.Context(), KeyProduct{}, NewProduct)
		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	})
}
