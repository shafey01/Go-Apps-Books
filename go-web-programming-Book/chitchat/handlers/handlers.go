package handlers

import (
	"log"
	"net/http"
	"regexp"
	"strconv"

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
	if r.Method == http.MethodPut {
		p.l.Println(r.URL.Path)

		req := regexp.MustCompile(`/([0-9+])`)
		regx := req.FindAllStringSubmatch(r.URL.Path, -1)
		if len(regx) != 1 {
			p.l.Println("Invalid URI more than one id")
			http.Error(w, "Invalid id", http.StatusBadRequest)
			return
		}
		if len(regx[0]) != 2 {

			p.l.Println("Invalid URI more than one capture")
			http.Error(w, "Invalid id", http.StatusBadRequest)
			return
		}
		idString := regx[0][1]
		id, err := strconv.Atoi(idString)
		if err != nil {

			p.l.Println("Invalid URI unable to convert to number")
			http.Error(w, "Invalid id", http.StatusBadRequest)

		}

		p.updateProduct(id, w, r)
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

func (p *Products) updateProduct(id int, w http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle PUT Request")
	NewProduct := &data.Product{}
	err := NewProduct.FromJSON(r.Body)
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
