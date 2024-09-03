package user

import (
	"phoenixia/internal/domain"
)

func (service *Service) GetAll() (users []domain.User, err error) {
	users, err = service.Repository.GetAll()
	return
}
