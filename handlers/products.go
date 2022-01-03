package handlers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/Drinnn/kool-products-ms/data"
	"github.com/gorilla/mux"
)

type Product struct {
	logger *log.Logger
}

func NewProduct(logger *log.Logger) *Product {
	return &Product{
		logger,
	}
}

func (p *Product) GetProducts(rw http.ResponseWriter, r *http.Request) {
	p.logger.Println("Handle GET Products")

	lp := data.GetProducts()

	err := lp.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
	}
}

func (p *Product) AddProduct(rw http.ResponseWriter, r *http.Request) {
	p.logger.Println("Handle POST Products")

	newProduct := &data.Product{}

	err := newProduct.FromJSON(r.Body)
	if err != nil {
		http.Error(rw, "Unable to unmarshal json", http.StatusBadRequest)
	}

	data.AddProduct(newProduct)
}

func (p *Product) UpdateProduct(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(rw, "Unable to convert id", http.StatusBadRequest)
		return
	}

	p.logger.Println("Handle PUT Products", id)

	updatedProduct := &data.Product{}

	err = updatedProduct.FromJSON(r.Body)
	if err != nil {
		http.Error(rw, "Unable to unmarshal json", http.StatusBadRequest)
	}

	err = data.UpdateProduct(id, updatedProduct)
	if err == data.ErrProductNotFound {
		http.Error(rw, "Product not found", http.StatusNotFound)
		return
	}

	if err != nil {
		http.Error(rw, "Product couldn't update", http.StatusInternalServerError)
		return
	}
}
