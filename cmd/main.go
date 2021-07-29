package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/Alptahta/personal-website-backend/internal/user/handlers"
	"github.com/Alptahta/personal-website-backend/internal/user/service"
)

func main() {
	l := log.New(os.Stdout, "product-api", log.LstdFlags)

	u := service.NewUser()
	mux := http.NewServeMux()
	handlers.NewUser(l, u).Register(mux)

	srv := &http.Server{
		Addr:         ":8080",
		Handler:      mux,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
