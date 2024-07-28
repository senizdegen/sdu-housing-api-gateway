package main

import (
	"github.com/julienschmidt/httprouter"
	"github.com/senizdegen/sdu-housing/api-gateway/internal/config"
	"github.com/senizdegen/sdu-housing/api-gateway/pkg/cache/freecache"
	"github.com/senizdegen/sdu-housing/api-gateway/pkg/jwt"
	"github.com/senizdegen/sdu-housing/api-gateway/pkg/logging"
)

func main() {
	logging.Init()
	logger := logging.GetLogger()
	logger.Println("logger initialized")

	logger.Println("config initializing")
	cfg := config.GetConfig()

	logger.Printf("cfg: %v", cfg)

	logger.Println("router intializing")
	router := httprouter.New()

	logger.Println("cache initializing")
	refreshTokenCache := freecache.NewCacheRepo(104857600)

	logger.Println("helpers initializing")
	jwtHelper := jwt.NewHelper(refreshTokenCache, logger)

	_ = router
	_ = jwtHelper
}
