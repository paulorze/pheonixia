package ollamaai

import (
	"bytes"
	"context"
	"strings"

	"github.com/google/uuid"
	"github.com/tmc/langchaingo/documentloaders"
	"github.com/tmc/langchaingo/textsplitter"
	"github.com/tmc/langchaingo/vectorstores/pgvector"
)

func (engine *Engine) PDFLoader(file []byte) (string, error) {
	reader := bytes.NewReader(file)
	size := int64(len(file))
	pdfloader := documentloaders.NewPDF(reader, size)

	ctx := context.Background()
	splitter := textsplitter.NewRecursiveCharacter() //! CAMBIAR ACA TAMANO DE DOCUMENTOS

	documents, err := pdfloader.LoadAndSplit(ctx, splitter)

	if err != nil {
		return "", err
	}

	tableName := strings.ReplaceAll(uuid.New().String(), "-", "")

	store, err := pgvector.New(
		ctx,
		pgvector.WithConnectionURL("postgres://postgres:postgres@localhost:5432/pgvector?sslmode=disable"), //! PASAR A VARIABLE DE ENTORNO
		pgvector.WithEmbeddingTableName(tableName),
		pgvector.WithEmbedder(engine.Embedder),
	)
	if err != nil {
		return "", err
	}

	_, errAdd := store.AddDocuments(ctx, documents)
	if errAdd != nil {
		return "", errAdd
	}

	return tableName, nil
}
