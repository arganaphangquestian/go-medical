package main

import (
	"github.com/arganaphangquestian/go-medical/blood"
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

	var r blood.Repository
	retry.ForeverSleep(2*time.Second, func(_ int) (err error) {
		r, err = blood.NewPostgres(cfg.DatabaseURL)
		if err != nil {
			log.Println(err)
		}
		return
	})
	defer r.Close()

	log.Println("Listening on port 8080...")
	s := blood.NewService(r)
	log.Fatal(blood.ListenGRPC(s, 8080))
}
