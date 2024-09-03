package user_test

import (
	customErrors "phoenixia/errors"
	"phoenixia/internal/domain"
	"phoenixia/internal/repositories/postgreSQL/user/mocks"
	"phoenixia/internal/services/user"
	"testing"

	"github.com/stretchr/testify/assert"
)

func StringPtr(str string) *string {
	return &str
}

func TestService_Create(t *testing.T) {

	tests := map[string]struct {
		user          domain.User
		err           error
		expectedError error
	}{
		"success": {
			user: domain.User{
				Username:  StringPtr("johndoe"),
				Password:  StringPtr("P4ssw0rd//"),
				Email:     StringPtr("johndoe@gmail.com"),
				FirstName: StringPtr("John"),
				LastName:  StringPtr("Doe"),
			},
			err:           nil,
			expectedError: nil,
		},
		"incomplete username": {
			user: domain.User{
				Username:  StringPtr(""),
				Password:  StringPtr("P4ssw0rd//"),
				Email:     StringPtr("johndoe@gmail.com"),
				FirstName: StringPtr("John"),
				LastName:  StringPtr("Doe"),
			},
			err:           &customErrors.IncompleteRegisterFields,
			expectedError: &customErrors.IncompleteRegisterFields,
		},
		"incomplete password": {
			user: domain.User{
				Username:  StringPtr("johndoe"),
				Password:  StringPtr(""),
				Email:     StringPtr("johndoe@gmail.com"),
				FirstName: StringPtr("John"),
				LastName:  StringPtr("Doe"),
			},
			err:           &customErrors.IncompleteRegisterFields,
			expectedError: &customErrors.IncompleteRegisterFields,
		},
		"incomplete email": {
			user: domain.User{
				Username:  StringPtr("johndoe"),
				Password:  StringPtr("P4ssw0rd//"),
				Email:     StringPtr(""),
				FirstName: StringPtr("John"),
				LastName:  StringPtr("Doe"),
			},
			err:           &customErrors.IncompleteRegisterFields,
			expectedError: &customErrors.IncompleteRegisterFields,
		},
		"incomplete firstname": {
			user: domain.User{
				Username:  StringPtr("johndoe"),
				Password:  StringPtr("P4ssw0rd//"),
				Email:     StringPtr("johndoe@gmail.com"),
				FirstName: StringPtr(""),
				LastName:  StringPtr("Doe"),
			},
			err:           &customErrors.IncompleteRegisterFields,
			expectedError: &customErrors.IncompleteRegisterFields,
		},
		"incomplete lastname": {
			user: domain.User{
				Username:  StringPtr("johndoe"),
				Password:  StringPtr("P4ssw0rd//"),
				Email:     StringPtr("johndoe@gmail.com"),
				FirstName: StringPtr("John"),
				LastName:  StringPtr(""),
			},
			err:           &customErrors.IncompleteRegisterFields,
			expectedError: &customErrors.IncompleteRegisterFields,
		},
		"invalid username length": {
			user: domain.User{
				Username:  StringPtr("doe"),
				Password:  StringPtr("P4ssw0rd//"),
				Email:     StringPtr("johndoe@gmail.com"),
				FirstName: StringPtr("John"),
				LastName:  StringPtr("Doe"),
			},
			err:           &customErrors.InvalidUsername,
			expectedError: &customErrors.InvalidUsername,
		},
		"invalid username character": {
			user: domain.User{
				Username:  StringPtr("johndoe/"),
				Password:  StringPtr("P4ssw0rd//"),
				Email:     StringPtr("johndoe@gmail.com"),
				FirstName: StringPtr("John"),
				LastName:  StringPtr("Doe"),
			},
			err:           &customErrors.InvalidUsername,
			expectedError: &customErrors.InvalidUsername,
		},
		"invalid password length": {
			user: domain.User{
				Username:  StringPtr("johndoe"),
				Password:  StringPtr("Pa0//"),
				Email:     StringPtr("johndoe@gmail.com"),
				FirstName: StringPtr("John"),
				LastName:  StringPtr("Doe"),
			},
			err:           &customErrors.InvalidPassword,
			expectedError: &customErrors.InvalidPassword,
		},
		"invalid password missing lowercase": {
			user: domain.User{
				Username:  StringPtr("johndoe"),
				Password:  StringPtr("PASSWORD40//"),
				Email:     StringPtr("johndoe@gmail.com"),
				FirstName: StringPtr("John"),
				LastName:  StringPtr("Doe"),
			},
			err:           &customErrors.InvalidPassword,
			expectedError: &customErrors.InvalidPassword,
		},
		"invalid password missing uppercase": {
			user: domain.User{
				Username:  StringPtr("johndoe"),
				Password:  StringPtr("password40//"),
				Email:     StringPtr("johndoe@gmail.com"),
				FirstName: StringPtr("John"),
				LastName:  StringPtr("Doe"),
			},
			err:           &customErrors.InvalidPassword,
			expectedError: &customErrors.InvalidPassword,
		},
		"invalid password missing number": {
			user: domain.User{
				Username:  StringPtr("johndoe"),
				Password:  StringPtr("Password//"),
				Email:     StringPtr("johndoe@gmail.com"),
				FirstName: StringPtr("John"),
				LastName:  StringPtr("Doe"),
			},
			err:           &customErrors.InvalidPassword,
			expectedError: &customErrors.InvalidPassword,
		},
		"invalid password missing special character": {
			user: domain.User{
				Username:  StringPtr("johndoe"),
				Password:  StringPtr("P4ssw0rd"),
				Email:     StringPtr("johndoe@gmail.com"),
				FirstName: StringPtr("John"),
				LastName:  StringPtr("Doe"),
			},
			err:           &customErrors.InvalidPassword,
			expectedError: &customErrors.InvalidPassword,
		},
		"invalid email missing username": {
			user: domain.User{
				Username:  StringPtr("johndoe"),
				Password:  StringPtr("P4ssw0rd//"),
				Email:     StringPtr("@gmail.com"),
				FirstName: StringPtr("John"),
				LastName:  StringPtr("Doe"),
			},
			err:           &customErrors.InvalidEmail,
			expectedError: &customErrors.InvalidEmail,
		},
		"invalid email invalid username": {
			user: domain.User{
				Username:  StringPtr("johndoe"),
				Password:  StringPtr("P4ssw0rd//"),
				Email:     StringPtr("jhon<h1>@gmail.com"),
				FirstName: StringPtr("John"),
				LastName:  StringPtr("Doe"),
			},
			err:           &customErrors.InvalidEmail,
			expectedError: &customErrors.InvalidEmail,
		},
		"invalid email missing @": {
			user: domain.User{
				Username:  StringPtr("johndoe"),
				Password:  StringPtr("P4ssw0rd//"),
				Email:     StringPtr("johndoegmail.com"),
				FirstName: StringPtr("John"),
				LastName:  StringPtr("Doe"),
			},
			err:           &customErrors.InvalidEmail,
			expectedError: &customErrors.InvalidEmail,
		},
		"invalid email missing mail server": {
			user: domain.User{
				Username:  StringPtr("johndoe"),
				Password:  StringPtr("P4ssw0rd//"),
				Email:     StringPtr("johndoe@.com"),
				FirstName: StringPtr("John"),
				LastName:  StringPtr("Doe"),
			},
			err:           &customErrors.InvalidEmail,
			expectedError: &customErrors.InvalidEmail,
		},
		"invalid email invalid mail server": {
			user: domain.User{
				Username:  StringPtr("johndoe"),
				Password:  StringPtr("P4ssw0rd//"),
				Email:     StringPtr("johndoe@gmail<h1>.com"),
				FirstName: StringPtr("John"),
				LastName:  StringPtr("Doe"),
			},
			err:           &customErrors.InvalidEmail,
			expectedError: &customErrors.InvalidEmail,
		},
		"invalid email missing domain": {
			user: domain.User{
				Username:  StringPtr("johndoe"),
				Password:  StringPtr("P4ssw0rd//"),
				Email:     StringPtr("john@gmail."),
				FirstName: StringPtr("John"),
				LastName:  StringPtr("Doe"),
			},
			err:           &customErrors.InvalidEmail,
			expectedError: &customErrors.InvalidEmail,
		},
		"invalid email invalid domain": {
			user: domain.User{
				Username:  StringPtr("johndoe"),
				Password:  StringPtr("P4ssw0rd//"),
				Email:     StringPtr("@gmail.co12"),
				FirstName: StringPtr("John"),
				LastName:  StringPtr("Doe"),
			},
			err:           &customErrors.InvalidEmail,
			expectedError: &customErrors.InvalidEmail,
		},
		"invalid first name length": {
			user: domain.User{
				Username:  StringPtr("johndoe"),
				Password:  StringPtr("P4ssw0rd//"),
				Email:     StringPtr("johndoe@gmail.com"),
				FirstName: StringPtr("J"),
				LastName:  StringPtr("Doe"),
			},
			err:           &customErrors.InvalidName,
			expectedError: &customErrors.InvalidName,
		},
		"invalid first name character": {
			user: domain.User{
				Username:  StringPtr("johndoe"),
				Password:  StringPtr("P4ssw0rd//"),
				Email:     StringPtr("johndoe@gmail.com"),
				FirstName: StringPtr("Jo/-.n"),
				LastName:  StringPtr("Doe"),
			},
			err:           &customErrors.InvalidName,
			expectedError: &customErrors.InvalidName,
		},
		"invalid last name length": {
			user: domain.User{
				Username:  StringPtr("johndoe"),
				Password:  StringPtr("P4ssw0rd//"),
				Email:     StringPtr("johndoe@gmail.com"),
				FirstName: StringPtr("John"),
				LastName:  StringPtr("D"),
			},
			err:           &customErrors.InvalidName,
			expectedError: &customErrors.InvalidName,
		},
		"invalid last name character": {
			user: domain.User{
				Username:  StringPtr("johndoe"),
				Password:  StringPtr("P4ssw0rd//"),
				Email:     StringPtr("johndoe@gmail.com"),
				FirstName: StringPtr("John"),
				LastName:  StringPtr("Do3"),
			},
			err:           &customErrors.InvalidName,
			expectedError: &customErrors.InvalidName,
		},
		"existing mail": {
			user: domain.User{
				Username:  StringPtr("johndoe"),
				Password:  StringPtr("P4ssw0rd//"),
				Email:     StringPtr("johndoe@gmail.com"),
				FirstName: StringPtr("John"),
				LastName:  StringPtr("Doe"),
			},
			err:           &customErrors.ExistingEntry,
			expectedError: &customErrors.ExistingEntry,
		},
		"invalid data": {
			user: domain.User{
				Username:  StringPtr("johndoe"),
				Password:  StringPtr("P4ssw0rd//"),
				Email:     StringPtr("johndoe@gmail.com"),
				FirstName: StringPtr("John"),
				LastName:  StringPtr("Doe"),
			},
			err:           &customErrors.InvalidEmail,
			expectedError: &customErrors.InvalidEmail,
		},
		"server error": {
			user: domain.User{
				Username:  StringPtr("johndoe"),
				Password:  StringPtr("P4ssw0rd//"),
				Email:     StringPtr("johndoe@gmail.com"),
				FirstName: StringPtr("John"),
				LastName:  StringPtr("Doe"),
			},
			err:           &customErrors.ServerError,
			expectedError: &customErrors.ServerError,
		},
	}

	for testname, subtest := range tests {
		t.Run(testname, func(t *testing.T) {
			mockUserRepository := new(mocks.MockUserRepository)
			userService := user.Service{Repository: mockUserRepository}
			mockUserRepository.On("Create").Return(subtest.err)
			err := userService.Create(subtest.user)

			assert.Equal(t, subtest.expectedError, err)
		})
	}

}
