package config

import (
	"calculate_product_packs/internal/domain"
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
	"strings"
)

type Config struct {
	PackSizes []domain.PackSize
	Port      string
}

func NewConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file, using default environment variables")
	}

	return &Config{
		PackSizes: getPackSizesFromEnv(),
		Port:      getPortFromEnv(),
	}
}

func getPackSizesFromEnv() []domain.PackSize {
	packSizesStr := os.Getenv("PACK_SIZES")
	if packSizesStr == "" {
		// Default pack sizes if not provided
		return []domain.PackSize{250, 500, 1000, 2000, 5000}
	}

	sizesStr := strings.Split(packSizesStr, ",")
	var sizes []domain.PackSize
	for _, s := range sizesStr {
		size, err := strconv.Atoi(strings.TrimSpace(s))
		if err == nil {
			sizes = append(sizes, domain.PackSize(size))
		}
	}

	if len(sizes) == 0 {
		// Fallback to default if parsing fails
		return []domain.PackSize{250, 500, 1000, 2000, 5000}
	}

	return sizes
}

func getPortFromEnv() string {
	port := os.Getenv("PORT")
	if port == "" {
		return "8080" // default port
	}
	return port
}
