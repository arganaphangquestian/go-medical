package main

import (
	"github.com/kelseyhightower/envconfig"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
)

type AppConfig struct {
	BloodURL   string `envconfig:"BLOOD_SERVICE_URL"`
	DiseaseURL string `envconfig:"DISEASE_SERVICE_URL"`
	GenderURL  string `envconfig:"GENDER_SERVICE_URL"`
	HistoryURL string `envconfig:"HISTORY_SERVICE_URL"`
	RoleURL    string `envconfig:"ROLE_SERVICE_URL"`
	UserURL    string `envconfig:"USER_SERVICE_URL"`
}

const defaultPort = "8080"

func main() {
	var cfg AppConfig
	err := envconfig.Process("", &cfg)
	if err != nil {
		log.Fatal(err)
	}
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	srv, err := NewGraphQLServer(cfg.BloodURL, cfg.DiseaseURL, cfg.GenderURL, cfg.HistoryURL, cfg.RoleURL, cfg.UserURL)
	if err != nil {
		log.Fatal(err)
	}

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", handler.NewDefaultServer(srv.ToExecutableSchema()))

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
