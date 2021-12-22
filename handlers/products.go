package handlers

import (
	"log"
	"net/http"
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

}
