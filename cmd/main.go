package main

import (
	"calculate_product_packs/internal/config"
	httphandler "calculate_product_packs/internal/interfaces/http"
	"calculate_product_packs/internal/repository"
	"calculate_product_packs/internal/usecases"
	"log"
	"net/http"
)

func main() {
	config := config.NewConfig()

	repo := repository.NewMemoryPackSizeRepository(config.PackSizes)

	useCase := usecases.NewCalculatePacksUseCase(repo)

	handler := httphandler.NewPackCalculatorHandler(useCase)
	router := httphandler.NewRouter(handler)

	log.Printf("Starting server on :%s\n", config.Port)
	log.Fatal(http.ListenAndServe(":"+config.Port, router))
}
