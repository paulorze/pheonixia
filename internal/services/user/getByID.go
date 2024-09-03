package user

import (
	customErrors "phoenixia/errors"
	"phoenixia/internal/domain"
)

func (service *Service) GetByID(id string) (user domain.User, err error) {
	if id == "" {
		err = &customErrors.InvalidId
		return
	}
	user, err = service.Repository.GetByID(id)
	return
}
