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

	model := os.Getenv("AZURE_OPENAI_RES")
	openaiApiUrl := os.Getenv("AZURE_OPENAI_ENDPOINT")
	openaiApiKey := os.Getenv("AZURE_OPENAI_KEY")
	//openaiApiVersion := os.Getenv("AZURE_OPENAI_VERSION")
	//embeddingModel := os.Getenv("AZURE_EMBEDDING_MODEL")

	llm, err := openai.New(
		openai.WithAPIType(openai.APITypeAzure),
		openai.WithModel(model),
		openai.WithBaseURL(openaiApiUrl),
		openai.WithToken(openaiApiKey),
		//openai.WithAPIVersion(openaiApiVersion),
		openai.WithEmbeddingModel("EMPTY"),
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
