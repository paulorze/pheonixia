package engine

func (service *Service) Ask(contexto, history, query string) (string, error) {
	response, err := service.Repository.Ask(contexto, history, query)
	return response, err
}
