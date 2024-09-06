package main

import (
	"errors"
	"fmt"
	"net"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"syscall"
	"time"

	"github.com/julienschmidt/httprouter"
	"github.com/senizdegen/sdu-housing/api-gateway/internal/client/property_service"
	"github.com/senizdegen/sdu-housing/api-gateway/internal/client/user_service"
	"github.com/senizdegen/sdu-housing/api-gateway/internal/config"
	"github.com/senizdegen/sdu-housing/api-gateway/internal/handlers/property"
	"github.com/senizdegen/sdu-housing/api-gateway/internal/handlers/users"
	"github.com/senizdegen/sdu-housing/api-gateway/pkg/handlers/metric"
	"github.com/senizdegen/sdu-housing/api-gateway/pkg/logging"
	"github.com/senizdegen/sdu-housing/api-gateway/pkg/shutdown"
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

	logger.Println("create and register handlers")

	metricHandler := metric.Handler{Logger: logger}
	metricHandler.Register(router)

	userService := user_service.NewService(cfg.UserService.URL, "/users", logger)
	usersHandler := users.Handler{UserService: userService, Logger: logger}
	usersHandler.Register(router)

	propertyService := property_service.NewService(cfg.PropertyService.URL, "/property", logger)

	propertyHandler := property.Handler{PropertyService: propertyService, Logger: logger}

	propertyHandler.Register(router)

	start(router, logger, cfg)
}

func start(router *httprouter.Router, logger logging.Logger, cfg *config.Config) {
	var server *http.Server
	var listener net.Listener

	if cfg.Listen.Type == "sock" {
		appDir, err := filepath.Abs(filepath.Dir(os.Args[0]))
		if err != nil {
			logger.Fatal(err)
		}
		socketPath := path.Join(appDir, "app.sock")
		logger.Infof("socket path: %s", socketPath)

		logger.Info("create and listen unix socket")
		listener, err = net.Listen("unix", socketPath)
		if err != nil {
			logger.Fatal(err)
		}
	} else {
		logger.Infof("bind application to host: %s and port: %s", cfg.Listen.BindIP, cfg.Listen.Port)

		var err error

		listener, err = net.Listen("tcp", fmt.Sprintf("%s:%s", cfg.Listen.BindIP, cfg.Listen.Port))
		if err != nil {
			logger.Fatal(err)
		}
	}

	server = &http.Server{
		Handler:      router,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	go shutdown.Graceful([]os.Signal{syscall.SIGABRT, syscall.SIGQUIT, syscall.SIGHUP, os.Interrupt, syscall.SIGTERM},
		server)

	logger.Println("application initialized and started")

	if err := server.Serve(listener); err != nil {
		switch {
		case errors.Is(err, http.ErrServerClosed):
			logger.Warn("server shutdown")
		default:
			logger.Fatal(err)
		}
	}
}
