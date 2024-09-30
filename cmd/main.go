package main

import (
	"calculate_product_packs/internal/config"
	httphandler "calculate_product_packs/internal/interfaces/http"
	"calculate_product_packs/internal/repository"
	"calculate_product_packs/internal/usecases"
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	// Initialize configuration
	config := config.NewConfig()

	// Set up repository and use cases
	repo := repository.NewMemoryPackSizeRepository(config.PackSizes)
	calculatePacksUseCase := usecases.NewCalculatePacksUseCase(repo)
	packSizesUseCase := usecases.NewPackSizesUseCase(repo)

	// Set up HTTP handler and router
	handler := httphandler.NewPackCalculatorHandler(calculatePacksUseCase, packSizesUseCase)
	router := httphandler.NewRouter(handler)

	// Configure the HTTP server
	srv := &http.Server{
		Addr:    ":" + config.Port,
		Handler: router,
	}

	// Channel to capture server errors
	serverErrors := make(chan error, 1)

	// Start the server in a goroutine
	go func() {
		log.Printf("Starting server on :%s\n", config.Port)
		serverErrors <- srv.ListenAndServe()
	}()

	// Set up graceful shutdown
	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt, syscall.SIGTERM)

	// Wait for server error or shutdown signal
	select {
	case err := <-serverErrors:
		log.Fatalf("Error starting server: %v", err)

	case <-shutdown:
		log.Println("Shutdown signal received")

		// Create a deadline for graceful shutdown
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		// Attempt graceful shutdown
		if err := srv.Shutdown(ctx); err != nil {
			log.Printf("Error during server shutdown: %v", err)
			// Force shutdown if graceful shutdown fails
			if err := srv.Close(); err != nil {
				log.Printf("Error closing server: %v", err)
			}
		}
	}

	log.Println("Server stopped")
}
