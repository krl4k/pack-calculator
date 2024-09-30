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
	config := config.NewConfig()

	repo := repository.NewMemoryPackSizeRepository(config.PackSizes)

	calculatePacksUseCase := usecases.NewCalculatePacksUseCase(repo)
	packSizesUseCase := usecases.NewPackSizesUseCase(repo)

	handler := httphandler.NewPackCalculatorHandler(calculatePacksUseCase, packSizesUseCase)
	router := httphandler.NewRouter(handler)

	srv := &http.Server{
		Addr:    ":" + config.Port,
		Handler: router,
	}

	serverErrors := make(chan error, 1)

	go func() {
		log.Printf("Starting server on :%s\n", config.Port)
		serverErrors <- srv.ListenAndServe()
	}()

	// Graceful shutdown
	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt, syscall.SIGTERM)

	select {
	case err := <-serverErrors:
		log.Fatalf("Error starting server: %v", err)

	case <-shutdown:
		log.Println("Shutdown signal received")

		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		if err := srv.Shutdown(ctx); err != nil {
			log.Printf("Error during server shutdown: %v", err)
			// Принудительное завершение, если graceful shutdown не удался
			if err := srv.Close(); err != nil {
				log.Printf("Error closing server: %v", err)
			}
		}
	}

	log.Println("Server stopped")
}
