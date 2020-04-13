package crud

import (
	"context"
	"fmt"
	"github.com/gorilla/mux"
	crud_config "github.com/sochoa/obsidian/crud/config"
	crud_object "github.com/sochoa/obsidian/crud/object"
	"log"
	"net/http"
	"os"
	"os/signal"
)

func Serve(cfg crud_config.ObjectStorageConfig) {
	requestRouter := mux.NewRouter()
	crud_object.SetupObjectRoutes(requestRouter, cfg)

	srv := &http.Server{
		Addr: cfg.BindPoint(),
		// Protect against Slowloris DOS attacks:  https://en.wikipedia.org/wiki/Slowloris_(computer_security)
		WriteTimeout: cfg.WriteTimeout,
		ReadTimeout:  cfg.ReadTimeout,
		IdleTimeout:  cfg.IdleTimeout,
		Handler:      requestLoggingMiddleware(requestRouter),
	}

	go func() {
		log.Println(fmt.Sprintf("Obsidian object storage service now listening on %s:%d with config:  %v", cfg.Host, cfg.Port, cfg))
		if err := srv.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c // Waiting for Interrupt

	ctx, cancel := context.WithTimeout(context.Background(), cfg.GracefulShutdownTimeout)
	defer cancel() // Why?
	srv.Shutdown(ctx)
	log.Println("Exiting...")
	os.Exit(0)
}
