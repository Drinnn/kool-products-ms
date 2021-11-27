package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/Drinnn/kool-products-ms/handlers"
)

func main() {
	logger := log.New(os.Stdout, "kool-products", log.LstdFlags)
	hh := handlers.NewHello(logger)
	gh := handlers.NewGoodbye(logger)

	sm := http.NewServeMux()
	sm.Handle("/", hh)
	sm.Handle("/goodbye", gh)

	s := http.Server{
		Addr:         ":9090",
		Handler:      sm,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	go func() {
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
