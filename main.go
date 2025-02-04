package main

import (
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

type Article struct {

}

func init() {
	if err := godotenv.Load(); err != nil {
		slog.Error("Could not read .env file", "error", err)
	}
}

func main() {
	url := fmt.Sprintf("https://newsdata.io/api/1/news?apikey=%s&country=es&language=es", os.Getenv("API_KEY"))
	res, err := http.Get(url)
	if err != nil {
		slog.Error("Error getting the content of www.google.es", "error", err)
	}

	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if res.StatusCode > 299 {
		slog.Error("Response failed", "status code", res.StatusCode, "body", body)
		os.Exit(1)
	}
	if err != nil {
		slog.Error("Error getting the content of www.google.es", "error", err)
		os.Exit(1)
	}
}
