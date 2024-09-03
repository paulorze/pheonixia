package user

import (
	customErrors "phoenixia/errors"
	"phoenixia/internal/domain"
	"phoenixia/utils"
)

func (service *Service) Create(user domain.User) (err error) {
	if *user.Username == "" || *user.Password == "" || *user.Email == "" || *user.FirstName == "" || *user.LastName == "" {
		return &customErrors.IncompleteRegisterFields
	}
	validatedUsername, err := utils.ValidateUsername(*user.Username)
	if err != nil {
		return err
	}
	validatedPassword, err := utils.ValidatePassword(*user.Password)
	if err != nil {
		return err
	}
	validatedEmail, err := utils.ValidateEmail(*user.Email)
	if err != nil {
		return err
	}
	validatedFirstName, err := utils.ValidateName(*user.FirstName)
	if err != nil {
		return err
	}
	validatedLastName, err := utils.ValidateName(*user.LastName)
	if err != nil {
		return err
	}
	role := "user"
	newUser := domain.User{
		Username:       &validatedUsername,
		Password:       &validatedPassword,
		Email:          &validatedEmail,
		FirstName:      &validatedFirstName,
		LastName:       &validatedLastName,
		Role:           &role,
		LastConnection: nil,
		Token:          nil,
	}
	err = service.Repository.Create(newUser)
	return
}
