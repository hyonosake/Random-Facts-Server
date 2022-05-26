package main

import (
	"context"
	"github.com/hyonosake/Random-Facts-Server/internal/server"
	"log"
)

func main() {

	ctx := context.Background()
	s, err := server.New(ctx)
	if err != nil {
		log.Fatalf("Unable to craete server: %v\n", err)
	}
	go s.HealthCheck(ctx)
	s.Run(ctx)

}
