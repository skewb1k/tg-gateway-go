package tggateway_test

import (
	"os"
	"testing"
	"time"

	tggateway "github.com/skewb1k/tg-gateway-go/v2"
)

var (
	phone  string
	token  string
	client *tggateway.Client
)

const requestInterval = 5 * time.Second

func TestMain(m *testing.M) {
	// Setup: fetch environment variables and create the client
	phone = os.Getenv("TEST_PHONE")
	if phone == "" {
		panic("phone env was not provided")
	}

	token = os.Getenv("TEST_API_TOKEN")
	if token == "" {
		panic("token env was not provided")
	}

	client = tggateway.NewClient(token)

	// Run the tests
	exitCode := m.Run()

	// Teardown (if needed)
	os.Exit(exitCode)
}
