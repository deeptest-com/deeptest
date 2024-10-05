package main

import (
	"context"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/llms/openai"
	"log"
	"os"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	model := os.Getenv("OPENAI_API_MODEL")
	openaiApiUrl := os.Getenv("OPENAI_API_URL")
	openaiApiKey := os.Getenv("OPENAI_API_KEY")

	llm, err := openai.New(
		openai.WithModel(model),
		openai.WithBaseURL(openaiApiUrl),
		openai.WithToken(openaiApiKey),
	)
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()
	completion, err := llms.GenerateFromSinglePrompt(ctx,
		llm,
		"The first man to walk on the moon",
		llms.WithTemperature(0.8),
		llms.WithStopWords([]string{"Armstrong"}),
	)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("The first man to walk on the moon:")
	fmt.Println(completion)
}
