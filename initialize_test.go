package aiozaiplatformgosdk

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	testApiKey string
	testClient *ClientV2
)

func init() {
	if err := loadEnvVariables(); err != nil {
		log.Fatalf("Failed to load environment variables: %v", err)
	}

	initializeClients()
}

func loadEnvVariables() error {
	if err := godotenv.Load(); err != nil {
		log.Printf("Warning: .env file not found, using system environment variables")
	}

	var missingVars []string

	testApiKey = os.Getenv("API_KEY")
	if testApiKey == "" {
		missingVars = append(missingVars, "API_KEY")
	}

	if len(missingVars) > 0 {
		return fmt.Errorf("missing required environment variables: %v", missingVars)
	}

	return nil
}

func initializeClients() {
	testClient = ClientBuilder(AuthCredentials{
		ApiKey: testApiKey,
	}).Build()
}
