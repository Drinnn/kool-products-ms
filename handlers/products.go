package handlers

import (
	"encoding/json"
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
	data, err := json.Marshal(lp)
	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
	}

	rw.Write(data)
}
