package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/thisisamr/SysWatch/config"
	"github.com/thisisamr/SysWatch/internal/metrics"
	"github.com/thisisamr/SysWatch/internal/server"
)

func main() {

	config.InitConfig()
	// The HTTP server with handler and address
	server := &http.Server{
		Addr:    fmt.Sprintf("%s:%s", config.Host, config.Port),
		Handler: server.NewServer(&metrics.Provider{}).Router,
	}
	// Server run context
	serverCtx, serverStopCtx := context.WithCancel(context.Background())

	// Listen for syscall signals for process to interrupt/quit
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	go func() {
		<-sig

		// Shutdown signal with grace period of 30 seconds
		shutdownCtx, cancel := context.WithTimeout(serverCtx, 30*time.Second)

		defer cancel() // This ensures resources are released after shutdown
		go func() {
			<-shutdownCtx.Done()
			if shutdownCtx.Err() == context.DeadlineExceeded {
				log.Fatal("graceful shutdown timed out.. forcing exit.")
			}
		}()

		// Trigger graceful shutdown
		err := server.Shutdown(shutdownCtx)
		if err != nil {
			log.Fatal(err)
		}
		serverStopCtx()
	}()

	// Run the server
	go func() {

		err := server.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			log.Fatal(err)
		}
	}()
	log.Printf("Server is running on http://%s:%s", config.Host, config.Port)
	// Wait for server context to be stopped
	<-serverCtx.Done()
	log.Println("Server stopped.")
}
