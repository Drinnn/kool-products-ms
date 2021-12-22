package handlers

import (
	"log"
	"net/http"

	"github.com/Drinnn/kool-products-ms/data"
)

type Product struct {
	logger *log.Logger
}

func NewProduct(logger *log.Logger) *Product {
	return &Product{
		logger,
	}
}

func (p *Product) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	lp := data.GetProducts()
	err := lp.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
	}
}
