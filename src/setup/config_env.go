package setup

import (
	"errors"

	"github.com/joho/godotenv"
)

// variable config
const developmentPath = ".env"
const testPath = "../../.env"

// LoadEnv ...
func LoadEnv(path string) error {
	err := godotenv.Load(path)
	if err != nil {
		err = errors.New("Error loading .env file")
	}
	return err
}

// LoadEnvTest ...
func LoadEnvTest() {
	LoadEnv(testPath)
}

// LoadEnvDevelopment ...
func LoadEnvDevelopment() {
	LoadEnv(developmentPath)
}
