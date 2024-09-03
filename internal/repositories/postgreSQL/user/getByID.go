package user

import (
	"errors"
	customErrors "phoenixia/errors"
	"phoenixia/internal/domain"

	"gorm.io/gorm"
)

func (repository *Repository) GetByID(id string) (user domain.User, err error) {
	err = repository.DB.Where("id = ?", id).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = customErrors.UserNotFound
			return
		}
		err = &customErrors.ServerError
		return
	}
	return
}
