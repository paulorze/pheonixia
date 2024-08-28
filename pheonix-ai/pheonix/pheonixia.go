package pheonix

import (
	"fmt"
	"github.com/tmc/langchaingo/chains"
	"github.com/tmc/langchaingo/llms/openai"
	"github.com/tmc/langchaingo/tools"
	"github.com/tmc/langchaingo/vectorstores"
)

type PhoenixIA struct {
	model     openai.LLM // Usa LLM para interactuar con el modelo de OpenAI
	retriever *vectorstores.VectorStore
	chain     *chains.Chain
}

func NewPhoenixIA(apiKey string) *PhoenixIA {
	// Configura el modelo de lenguaje OpenAI
	model := openai.LLM{
		APIKey: apiKey,
		Model:  "gpt-3.5-turbo", // Ajusta el modelo según sea necesario
	}

	return &PhoenixIA{
		model: model,
	}
}

func (ai *PhoenixIA) Ingest(pdfFilePath string) error {
	// Extrae el texto del archivo PDF
	docs, err := tools.ExtractTextFromFile(pdfFilePath) // Verifica si esta función existe
	if err != nil {
		return fmt.Errorf("error al extraer texto del PDF: %v", err)
	}

	// Crea un vector store a partir de los documentos
	vectorStore, err := vectorstores.NewFromDocuments(docs) // Verifica si esta función existe
	if err != nil {
		return fmt.Errorf("error al crear el vector store: %v", err)
	}

	ai.retriever = vectorStore
	ai.chain = chains.NewChain(ai.retriever, ai.model) // Verifica cómo se crea una cadena en la biblioteca

	return nil
}

func (ai *PhoenixIA) Ask(query string) (string, error) {
	if ai.chain == nil {
		return "", fmt.Errorf("debes ingresar un PDF primero")
	}

	// Utiliza la cadena para responder la consulta
	response, err := ai.chain.Run(query) // Verifica el método correcto en la cadena
	if err != nil {
		return "", err
	}
	return response, nil
}

func (ai *PhoenixIA) Clear() {
	ai.retriever = nil
	ai.chain = nil
}
