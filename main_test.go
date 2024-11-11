package evolution_test

import (
	"os"
	"testing"

	"github.com/joho/godotenv"
	"github.com/verbeux-ai/evolution-client-go"
)

var client *evolution.Client

func TestMain(m *testing.M) {
	_ = godotenv.Load(".env")

	client = evolution.NewClient(
		evolution.WithApiKey(os.Getenv("API_KEY")),
		evolution.WithBaseUrl(os.Getenv("BASE_URL")),
	)

	os.Exit(m.Run())
}
