package data

import (
	"encoding/json"
	"fmt"
	"io"
	"time"
)

type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float32 `json:"price"`
	SKU         string  `json:"sku"`
	CreatedOn   string  `json:"-"`
	UpdatedOn   string  `json:"-"`
	DeletedOn   string  `json:"-"`
}

var ErrProductNotFound = fmt.Errorf("Product Not Found")

// Decode data
func (p *Product) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(p)
}

// list of Products
type Products []*Product

// Encode data
func (p *Products) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(p)
}

func GetProducts() Products {
	return productList
}

func AddProduct(p *Product) {
	id := getNextId()
	p.ID = id

	productList = append(productList, p)

}

func UpdateProduct(id int, p *Product) error {

	_, pos, err := findProduct(id)

	if err != nil {
		return err
	}
	p.ID = id
	productList[pos] = p
	return nil
}

// Utilits functions
func findProduct(id int) (*Product, int, error) {
	for i, v := range productList {
		if v.ID == id {
			return v, i, nil
		}
	}
	return nil, -1, ErrProductNotFound
}

func getNextId() int {

	Last := productList[len(productList)-1]
	return Last.ID + 1
}

var productList = []*Product{
	&Product{
		ID:          1,
		Name:        "Latte",
		Description: "Frothy milky coffee",
		Price:       2.45,
		SKU:         "abc323",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},

	&Product{
		ID:          2,
		Name:        "Espresso",
		Description: "Short and strong coffee without milk",
		Price:       1.99,
		SKU:         "fjd34",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
}
