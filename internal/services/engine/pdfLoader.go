package engine

func (service *Service) PDFLoader(file []byte) (string, error) {
	tablename, err := service.Repository.PDFLoader(file)

	return tablename, err
}
