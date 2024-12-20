package data

import (
	"encoding/json"
	"fmt"
	"io"
	"regexp"
	"time"

	"github.com/go-playground/validator/v10"
)

type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name" validate:"required"`
	Description string  `json:"description"`
	Price       float32 `json:"price"`
	SKU         string  `json:"sku" validate:"required,sku"`
	CreatedOn   string  `json:"-"`
	UpdatedOn   string  `json:"-"`
	DeletedOn   string  `json:"-"`
}

var ErrProductNotFound = fmt.Errorf("Product Not Found")
var ErrEmptyList = fmt.Errorf("Empty List")

// Decode data
func (p *Product) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(p)
}

// Validation function
func (p *Product) Validate() error {
	validate := validator.New()
	validate.RegisterValidation("sku", validateSKU)
	return validate.Struct(p)
}

// Validate SKU
func validateSKU(fl validator.FieldLevel) bool {
	reg := regexp.MustCompile(`[a-z]+-[a-z]+-[a-z]+`)
	matches := reg.FindAllString(fl.Field().String(), -1)
	if len(matches) != 1 {
		return false
	}
	return true
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

	pos := findProductByID(id)

	if pos == -1 {
		return ErrProductNotFound
	}
	p.ID = id
	productList[pos] = p
	return nil
}
func DelelteProduct(id int) error {

	pos := findProductByID(id)
	if pos == -1 {
		return ErrProductNotFound

	}
	if len(productList) >= 1 {
		productList = append(productList[:pos], productList[pos+1:]...)
	}

	if len(productList) == 0 {
		return ErrEmptyList
	}
	return nil
}

// Utilits functions
func findProductByID(id int) int {
	for i, v := range productList {
		if v.ID == id {
			return i
		}
	}
	return -1
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
