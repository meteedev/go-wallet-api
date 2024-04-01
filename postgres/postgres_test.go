package postgres

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateDatabaseUrl(t *testing.T) {
	// Set up environment variables for testing
	os.Setenv("POSTGRES_HOST", "localhost")
	os.Setenv("POSTGRES_PORT", "5432")
	os.Setenv("POSTGRES_USER", "testuser")
	os.Setenv("POSTGRES_PASSWORD", "testpassword")
	os.Setenv("POSTGRES_DB_NAME", "testdb")
	os.Setenv("POSTGRES_SSL_MODE", "disable")

	expectedURL := "host=localhost port=5432 user=testuser password=testpassword dbname=testdb sslmode=disable"

	// Call the function and check if the generated URL matches the expected URL
	actualURL := generateDatabaseUrl()
	assert.Equal(t, expectedURL, actualURL)
}

func TestInitDbConfiguration(t *testing.T) {
	// Set up environment variables for testing
	os.Setenv("POSTGRES_HOST", "localhost")
	os.Setenv("POSTGRES_PORT", "5432")
	os.Setenv("POSTGRES_USER", "testuser")
	os.Setenv("POSTGRES_PASSWORD", "testpassword")
	os.Setenv("POSTGRES_DB_NAME", "testdb")
	os.Setenv("POSTGRES_SSL_MODE", "disable")

	// Call the function and check if the returned configuration matches the expected values
	expectedConfig := &DbConfiguration{
		Host:     "localhost",
		Port:     "5432",
		User:     "testuser",
		Password: "testpassword",
		DbName:   "testdb",
		SslMode:  "disable",
	}

	actualConfig := initDbConfiguration()
	assert.Equal(t, expectedConfig, actualConfig)
}
