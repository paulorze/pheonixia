package engine

func (service *Service) DocumentsRetriever(tablename, query string) (string, error) {
	documents, err := service.Repository.DocumentsRetriever(tablename, query)

	return documents, err
}
