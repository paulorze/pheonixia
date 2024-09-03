package engine

import (
	"log"
	"phoenixia/internal/repositories/engine/ollamaai"

	"github.com/tmc/langchaingo/embeddings"
	"github.com/tmc/langchaingo/llms/ollama"
)

func NewOllamaAI() (*ollamaai.Engine, error) {
	modelinstance, err := ollama.New(ollama.WithModel("llama3.1"))
	if err != nil {
		return nil, err
	}

	embedderinstance, err := embeddings.NewEmbedder(modelinstance)
	if err != nil {
		log.Fatal(err)
	}

	return &ollamaai.Engine{
		Model:    modelinstance,
		Embedder: embedderinstance,
	}, nil
}
