package main

import (
	"log"
	"net/http"
	"os"

	"github.com/co0p/go-tls-watch/pkg/integration/slack"
	"github.com/co0p/go-tls-watch/pkg/integration/web"

	"github.com/co0p/go-tls-watch/pkg/usecases"
)

func main() {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "9991"
	}

	slackApiKey := os.Getenv("SLACK_API_KEY")
	if len(slackApiKey) == 0 {
		log.Fatal("missing env var SLACK_API_KEY")
	}

	fetchClient := web.NewFetchClient()
	slackClient := slack.NewSlackClient(slackApiKey)

	validateUsecase := usecases.ValidateUsecase{Client: &fetchClient}

	webHandler := web.Handler{ValidateUsecase: &validateUsecase}
	slackHandler := slack.Handler{
		ValidateUsecase: &validateUsecase,
		Client:          slackClient,
	}

	mux := http.NewServeMux()
	mux.Handle("/api/validate", webHandler.Handle())
	go slackHandler.Handle()

	log.Printf("starting app on port %s ...", port)
	if err := http.ListenAndServe(":"+port, mux); err != nil {
		log.Fatalf("failed to start app: %s", err.Error())
	}
}
