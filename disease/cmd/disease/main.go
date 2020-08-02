package main

import (
	"github.com/arganaphangquestian/go-medical/disease"
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

	var r disease.Repository
	retry.ForeverSleep(2*time.Second, func(_ int) (err error) {
		r, err = disease.NewElasticRepository(cfg.DatabaseURL)
		if err != nil {
			log.Println(err)
		}
		return
	})
	defer r.Close()

	log.Println("Listening on port 8080...")
	s := disease.NewService(r)
	log.Fatal(disease.ListenGRPC(s, 8080))
}
