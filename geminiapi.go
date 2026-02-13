package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv" // 1. Import the loader
	"google.golang.org/genai"
)

func main() {
    // Load .env file
    err := godotenv.Load()
    if err != nil {
        log.Fatal("Error loading .env file")
    }

    ctx := context.Background()

    // Pass the API Key directly in the config
    client, err := genai.NewClient(ctx, &genai.ClientConfig{
        APIKey: os.Getenv("GOOGLE_API_KEY"),
    })
    if err != nil {
        log.Fatal(err)
    }

    // Using gemini-2.0-flash as it is the current standard stable ID
    result, err := client.Models.GenerateContent(
        ctx,
        "gemini-2.5-flash", 
        genai.Text("what is the ai bubble"),
        nil,
    )
    if err != nil {
        log.Fatal(err)
    }

    // The SDK's result.Text() helper is usually sufficient
    fmt.Println(result.Text())
}