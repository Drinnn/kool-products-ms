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
	if r.Method == http.MethodGet {
		p.getProducts(rw, r)
		return
	}

	if r.Method == http.MethodPost {
		p.addProduct(rw, r)
		return
	}

	rw.WriteHeader(http.StatusMethodNotAllowed)
}

func (p *Product) getProducts(rw http.ResponseWriter, r *http.Request) {
	p.logger.Println("Handle GET Products")

	lp := data.GetProducts()

	err := lp.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
	}
}

func (p *Product) addProduct(rw http.ResponseWriter, r *http.Request) {
	p.logger.Println("Handle POST Products")

	newProduct := &data.Product{}

	err := newProduct.FromJSON(r.Body)
	if err != nil {
		http.Error(rw, "Unable to unmarshal json", http.StatusBadRequest)
	}

	data.AddProduct(newProduct)
}
