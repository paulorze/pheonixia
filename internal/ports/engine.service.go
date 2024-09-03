package ports

type EngineService interface {
	Ask(contexto, history, query string) (string, error)
	DocumentsRetriever(tablename, query string) (string, error)
	PDFLoader(file []byte) (string, error)
}
