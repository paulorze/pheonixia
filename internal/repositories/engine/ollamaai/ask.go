package ollamaai

import (
	"context"

	"github.com/tmc/langchaingo/llms"
)

func (engine *Engine) Ask(contexto, history, query string) (string, error) {
	ctx := context.Background()

	systemContent := "Eres un asistente virtual para empleados de servicio al cliente. Contesta la consulta de manera concisa pero clara. Si la respuesta no se encuentra en el contexto provisto, simplemente dí que no lo sabes y no contestes bajo ningún concepto. No aclares de qué documento proviene la respuesta. Ten en cuenta solamente el contenido del siguiente contexto e historial de conversación:"

	systemContext := "Contexto: " + contexto

	systemHistory := "Historial de Conversación: " + history

	systemQuery := "Consulta: " + query

	content := []llms.MessageContent{
		llms.TextParts(llms.ChatMessageTypeSystem, systemContent),
		llms.TextParts(llms.ChatMessageTypeSystem, systemContext),
		llms.TextParts(llms.ChatMessageTypeSystem, systemHistory),
		llms.TextParts(llms.ChatMessageTypeHuman, systemQuery),
	}

	var response string

	completion, err := engine.Model.GenerateContent(
		ctx,
		content,
		llms.WithStreamingFunc(func(ctx context.Context, chunk []byte) error {
			response += (string(chunk))
			return nil
		}),
	)

	if err != nil {
		return "", err
	}
	_ = completion

	return response, nil

}
