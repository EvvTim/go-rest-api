package main

import (
	"context"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"go-rest-api/internal/config"
	"go-rest-api/internal/user"
	"go-rest-api/pkg/client/mongodb"
	"go-rest-api/pkg/logging"
	"net"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"time"
)

const (
	socketName = "app.sock"
	socket     = "sock"
)

func main() {
	logger := logging.GetLogger()
	logger.Info("create router")
	router := httprouter.New()

	cfg := config.GetConfig()

	cfgMongoDB := cfg.MongoDB

	mongoDBClient, err := mongodb.NewClient(
		context.Background(),
		cfgMongoDB.Host,
		cfgMongoDB.Port,
		cfgMongoDB.Username,
		cfgMongoDB.Password,
		cfgMongoDB.Database,
		cfgMongoDB.AuthDB,
	)

	if err != nil {
		panic(err)
	}

	fmt.Println(mongoDBClient)

	//storage := db.NewStorage(mongoDBClient, cfg.MongoDB.Collection, logger)

	logger.Info("register user handler")
	handler := user.NewHandler(logger)
	handler.Register(router)

	start(router, cfg)
}

func start(router *httprouter.Router, cfg *config.Config) {
	logger := logging.GetLogger()
	logger.Info("start application")

	var listener net.Listener
	var listenErr error

	if cfg.Listen.Type == socket {
		logger.Info("detect app path")
		appDir, err := filepath.Abs(filepath.Dir(os.Args[0]))
		if err != nil {
			logger.Fatal(err)
		}

		logger.Info("create socket")
		socketPath := path.Join(appDir, socketName)

		logger.Info("listen unix socket")
		listener, listenErr = net.Listen("unix", socketPath)
		logger.Infof("server is listening unix socket: %s", socketPath)
	} else {
		logger.Info("listen tcp")
		listener, listenErr = net.Listen("tcp", fmt.Sprintf("%s:%s",
			cfg.Listen.BindIP,
			cfg.Listen.Port,
		))
	}

	if listenErr != nil {
		logger.Fatal(listenErr)
	}

	server := &http.Server{
		Handler:      router,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	logger.Infof("server is listening port %s:%s", cfg.Listen.BindIP, cfg.Listen.Port)

	logger.Fatal(server.Serve(listener))
}
