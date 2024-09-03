package ollamaai

import (
	"phoenixia/internal/ports"

	"github.com/tmc/langchaingo/embeddings"
	"github.com/tmc/langchaingo/llms/ollama"
)

var _ ports.EngineRepository = &Engine{}

type Engine struct {
	Model    *ollama.LLM
	Embedder embeddings.Embedder
}
