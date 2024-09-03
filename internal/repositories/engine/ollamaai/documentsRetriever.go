package ollamaai

import (
	"context"
	"log"
	"strconv"

	"github.com/tmc/langchaingo/vectorstores"
	"github.com/tmc/langchaingo/vectorstores/pgvector"
)

func (engine *Engine) DocumentsRetriever(tableName, query string) (string, error) {
	ctx := context.Background()
	store, err := pgvector.New(
		ctx,
		pgvector.WithConnectionURL("postgres://postgres:postgres@localhost:5432/pgvector?sslmode=disable"), //! PASAR A VARIABLE DE ENTORNO
		pgvector.WithEmbeddingTableName(tableName),
		pgvector.WithEmbedder(engine.Embedder),
	)
	if err != nil {
		return "", err
	}

	docs, err := store.SimilaritySearch(
		ctx,
		query,
		5,                                     //! CAMBIAR ACA LA CANTIDAD DE DOCUMENTOS
		vectorstores.WithScoreThreshold(0.05), //! CAMBIAR ACA EL LIMITE DE COINCIDENCIA PARA RESULTADOS
	)

	if err != nil {
		log.Fatal(err)
		return "", err
	}

	var content string
	for i, doc := range docs {
		content = content + "Documento N° " + strconv.Itoa(i) + " :"
		content = content + doc.PageContent + "\n"
	}

	return content, nil
}
