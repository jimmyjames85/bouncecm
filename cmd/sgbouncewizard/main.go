package main

import (
	"log"

	"github.com/jimmyjames85/bouncecm/internal/config"
	"github.com/jimmyjames85/bouncecm/internal/sgbouncewizard"
)

func main() {

	log.Println("the Server has Started")
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Error reading config", err)
	}

	srv, err := sgbouncewizard.NewServer(cfg)

	if err != nil {
		log.Fatalf("Error Starting Server:", err)
		return
	}
	srv.Serve(cfg.APIPort)
}
