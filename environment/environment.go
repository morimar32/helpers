package env

import (
	"fmt"
	"os"
	"strings"

	"github.com/morimar32/helpers/encryption"

	"github.com/joho/godotenv"
)

// LoadEnvironmentFile loads the appropriate .env config file, based on the defined ENV environment variable
func LoadEnvironmentFile() error {
	env := os.Getenv("ENV")
	if len(env) <= 0 {
		env = "Dev"
	}
	envPath := strings.ToLower(fmt.Sprintf(".%s.env", env))
	err := godotenv.Load(envPath)
	if err != nil {
		return err
	}
	return nil
}

// GetValue retrieves a value from the config for the given key, or an error if the key is not found
func GetValue(key string) (string, error) {
	value := os.Getenv(key)
	if len(value) <= 0 {
		return "", fmt.Errorf("could not find value for '%s'", key)
	}
	return value, nil
}

// GetValueWithDefault retrieves a value from the config for the given key, or the default if the key is not found
func GetValueWithDefault(key string, defaultVal string) string {
	value, err := GetValue(key)
	if err != nil {
		return defaultVal
	}
	return value
}

// GetEncryptedValue retrieves a value from the config for the given key and decrypts is, or an error if the key is not found
func GetEncryptedValue(key string) (string, error) {
	encValue, err := GetValue(key)
	if err != nil {
		return "", err
	}

	decValue, err := encryption.Decrypt(encValue)
	if err != nil {
		return "", err
	}
	if len(decValue) <= 0 {
		return "", fmt.Errorf("Could not decrypt value for '%s", key)
	}
	return decValue, nil
}

// GetEncryptedValueWithDefault retrieves a value from the config for the given key and decrypts it, or the default if the key is not found
func GetEncryptedValueWithDefault(key string, defaultVal string) string {
	value, err := GetEncryptedValue(key)
	if err != nil {
		return defaultVal
	}
	return value
}
