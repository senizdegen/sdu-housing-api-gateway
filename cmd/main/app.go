package main

import (
	"github.com/senizdegen/sdu-housing/api-gateway/internal/config"
	"github.com/senizdegen/sdu-housing/api-gateway/pkg/logging"
)

func main() {
	logging.Init()
	logger := logging.GetLogger()
	logger.Println("logger initialized")

	logger.Println("config initializing")
	cfg := config.GetConfig()

	logger.Printf("cfg: %v", cfg)
}
