package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"
	"github.com/nicholasjackson/env"

	"github.com/Drinnn/kool-products-ms/handlers"
)

var bindAddress = env.String("BIND_ADDRESS", false, ":9090", "Bind address for the server")

func main() {
	env.Parse()

	logger := log.New(os.Stdout, "kool-products", log.LstdFlags)

	ph := handlers.NewProduct(logger)

	sm := mux.NewRouter()

	getRouter := sm.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc("/", ph.GetProducts)

	putRouter := sm.Methods(http.MethodPut).Subrouter()
	putRouter.HandleFunc("/{id:[0-9]+}", ph.UpdateProduct)

	postRouter := sm.Methods(http.MethodPost).Subrouter()
	postRouter.HandleFunc("/", ph.AddProduct)

	// sm.Handle("/products", ph)

	s := http.Server{
		Addr:         *bindAddress,
		Handler:      sm,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	go func() {
		logger.Println("Server running on port", *bindAddress)

		err := s.ListenAndServe()
		if err != nil {
			log.Fatal(err)
		}
	}()

	signChan := make(chan os.Signal)
	signal.Notify(signChan, os.Interrupt)
	signal.Notify(signChan, os.Kill)

	sig := <-signChan
	logger.Println("Received terminate, graceful shutdown", sig)

	tc, _ := context.WithTimeout(context.Background(), 30*time.Second)

	s.Shutdown(tc)
}
