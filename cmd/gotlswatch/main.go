package main

import (
	"log"
	"net/http"
	"os"

	"github.com/co0p/go-tls-watch/pkg/integration/web"

	"github.com/co0p/go-tls-watch/pkg/usecases"
)

func main() {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "9991"
	}

	fetchClient := web.FetchClient{}

	validateUsecase := usecases.ValidateUsecase{Client: &fetchClient}

	webHandler := web.Handler{ValidateUsecase: &validateUsecase}

	mux := http.NewServeMux()
	mux.Handle("/api/validate", webHandler.Handle())

	log.Printf("starting app on port %s ...", port)
	if err := http.ListenAndServe(":"+port, mux); err != nil {
		log.Fatalf("failed to start app: %s", err.Error())
	}
}
