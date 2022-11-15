package infrastructure

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/moriuriel/go-task-api/infrastructure/routes"
)

type HTTPServer struct {
	routes *routes.Routes
}

func NewHTTPServer() *HTTPServer {
	return &HTTPServer{
		routes: routes.NewRoutes(),
	}
}

func (a HTTPServer) Start() {
	server := &http.Server{
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		Addr:         fmt.Sprintf(":%s", os.Getenv("PORT")),
		Handler:      a.routes.BuildRoutes(),
	}

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		log.Println("Starting HTTP Server in port:", os.Getenv("PORT"))
		log.Fatal(server.ListenAndServe())
	}()

	<-stop

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer func() {
		cancel()
	}()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown Failed")
	}

	log.Println("Service down")
}
