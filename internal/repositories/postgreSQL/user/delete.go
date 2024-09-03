package user

import (
	customErrors "phoenixia/errors"
	"phoenixia/internal/domain"
)

func (userDao *Repository) Delete(id string) (err error) {
	result := userDao.DB.Delete(&domain.User{}, id)
	if result.Error != nil {
		err = &customErrors.ServerError
		return
	}
	if result.RowsAffected == 0 {
		err = customErrors.UserNotFound
		return
	}
	return
}
