package user

import (
	customErrors "phoenixia/errors"
	"phoenixia/internal/domain"
)

func (repository *Repository) GetAll() (users []domain.User, err error) {
	err = repository.DB.Find(&users).Error
	if err != nil {
		err = &customErrors.ServerError
		return
	}
	return
}
