package main

import (
	"fmt"
	"github.com/EvvTim/go-rest-api/internal/user"
	"github.com/julienschmidt/httprouter"
	"log"
	"net"
	"net/http"
	"time"
)

func main() {
	fmt.Println("Create router")
	router := httprouter.New()

	handler := user.NewHandler()

	handler.Register(router)

	start(router)
}

func start(router *httprouter.Router) {
	listener, err := net.Listen("tcp", "127.0.0.1:8080")
	if err != nil {
		panic(err)
	}

	server := &http.Server{
		Handler:      router,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
	}

	log.Fatal(server.Serve(listener))
}
