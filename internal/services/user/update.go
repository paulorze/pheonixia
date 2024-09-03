package user

import (
	customErrors "phoenixia/errors"
	"phoenixia/internal/domain"
	"phoenixia/utils"
)

func (service *Service) Update(updatedUser domain.User) (err error) {
	if updatedUser.ID == 0 {
		err = &customErrors.InvalidId
		return
	}
	validatedEmail, validationErr := utils.ValidateEmail(*updatedUser.Email)
	if validationErr != nil {
		err = validationErr
		return
	}
	updatedUser.Email = &validatedEmail
	validatedFirstName, validationErr := utils.ValidateName(*updatedUser.FirstName)
	if validationErr != nil {
		err = validationErr
		return
	}
	updatedUser.FirstName = &validatedFirstName
	validatedLastName, validationErr := utils.ValidateName(*updatedUser.LastName)
	if validationErr != nil {
		err = validationErr
		return
	}
	updatedUser.LastName = &validatedLastName

	err = service.Repository.Update(updatedUser)
	return
}
