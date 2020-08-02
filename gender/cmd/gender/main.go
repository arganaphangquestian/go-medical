package main

import (
	"github.com/arganaphangquestian/go-medical/gender"
	"github.com/kelseyhightower/envconfig"
	"github.com/tinrab/kit/retry"
	"log"
	"time"
)

type Config struct {
	DatabaseURL string `envconfig:"DATABASE_URL"`
}

func main() {
	var cfg Config
	err := envconfig.Process("", &cfg)
	if err != nil {
		log.Fatal(err)
	}

	var r gender.Repository
	retry.ForeverSleep(2*time.Second, func(_ int) (err error) {
		r, err = gender.NewPostgres(cfg.DatabaseURL)
		if err != nil {
			log.Println(err)
		}
		return
	})
	defer r.Close()

	log.Println("Listening on port 8080...")
	s := gender.NewService(r)
	log.Fatal(gender.ListenGRPC(s, 8080))
}
