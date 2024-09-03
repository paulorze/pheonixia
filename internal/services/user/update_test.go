package user_test

import (
	customErrors "phoenixia/errors"
	"phoenixia/internal/domain"
	"phoenixia/internal/repositories/postgreSQL/user/mocks"
	"phoenixia/internal/services/user"
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func TestService_Update(t *testing.T) {

	tests := map[string]struct {
		user          domain.User
		err           error
		expectedError error
	}{
		"success": {
			user: domain.User{
				Model: gorm.Model{
					ID: 1,
				},
				Username:  StringPtr("janedoe"),
				Password:  StringPtr("P4ssw0rd//"),
				Email:     StringPtr("janedoe2@gmail.com"),
				FirstName: StringPtr("Janet"),
				LastName:  StringPtr("Doensky"),
			},
			err:           nil,
			expectedError: nil,
		},
		"invalid id": {
			user: domain.User{
				Model: gorm.Model{
					ID: 0,
				},
				Username:  StringPtr("janedoe"),
				Password:  StringPtr("P4ssw0rd//"),
				Email:     StringPtr("janedoe2@gmail.com"),
				FirstName: StringPtr("Janet"),
				LastName:  StringPtr("Doensky"),
			},
			err:           &customErrors.InvalidId,
			expectedError: &customErrors.InvalidId,
		},
		"invalid email missing username": {
			user: domain.User{
				Model: gorm.Model{
					ID: 1,
				},
				Username:  StringPtr("janedoe"),
				Password:  StringPtr("P4ssw0rd//"),
				Email:     StringPtr("@gmail.com"),
				FirstName: StringPtr("Janet"),
				LastName:  StringPtr("Doensky"),
			},
			err:           &customErrors.InvalidEmail,
			expectedError: &customErrors.InvalidEmail,
		},
		"invalid email invalid username": {
			user: domain.User{
				Model: gorm.Model{
					ID: 1,
				},
				Username:  StringPtr("janedoe"),
				Password:  StringPtr("P4ssw0rd//"),
				Email:     StringPtr("janedoe<h1>@gmail.com"),
				FirstName: StringPtr("Janet"),
				LastName:  StringPtr("Doensky"),
			},
			err:           &customErrors.InvalidEmail,
			expectedError: &customErrors.InvalidEmail,
		},
		"invalid email missing @": {
			user: domain.User{
				Model: gorm.Model{
					ID: 1,
				},
				Username:  StringPtr("janedoe"),
				Password:  StringPtr("P4ssw0rd//"),
				Email:     StringPtr("janedoe3gmail.com"),
				FirstName: StringPtr("Janet"),
				LastName:  StringPtr("Doensky"),
			},
			err:           &customErrors.InvalidEmail,
			expectedError: &customErrors.InvalidEmail,
		},
		"invalid email missing mail server": {
			user: domain.User{
				Model: gorm.Model{
					ID: 1,
				},
				Username:  StringPtr("janedoe"),
				Password:  StringPtr("P4ssw0rd//"),
				Email:     StringPtr("janedoe3@.com"),
				FirstName: StringPtr("Janet"),
				LastName:  StringPtr("Doensky"),
			},
			err:           &customErrors.InvalidEmail,
			expectedError: &customErrors.InvalidEmail,
		},
		"invalid email invalid mail server": {
			user: domain.User{
				Model: gorm.Model{
					ID: 1,
				},
				Username:  StringPtr("janedoe"),
				Password:  StringPtr("P4ssw0rd//"),
				Email:     StringPtr("janedoe3@gmail<h1>.com"),
				FirstName: StringPtr("Janet"),
				LastName:  StringPtr("Doensky"),
			},
			err:           &customErrors.InvalidEmail,
			expectedError: &customErrors.InvalidEmail,
		},
		"invalid email missing domain": {
			user: domain.User{
				Model: gorm.Model{
					ID: 1,
				},
				Username:  StringPtr("janedoe"),
				Password:  StringPtr("P4ssw0rd//"),
				Email:     StringPtr("janedoe3@gmail."),
				FirstName: StringPtr("Janet"),
				LastName:  StringPtr("Doensky"),
			},
			err:           &customErrors.InvalidEmail,
			expectedError: &customErrors.InvalidEmail,
		},
		"invalid email invalid domain": {
			user: domain.User{
				Model: gorm.Model{
					ID: 1,
				},
				Username:  StringPtr("janedoe"),
				Password:  StringPtr("P4ssw0rd//"),
				Email:     StringPtr("janedoe3@gmail.ar12"),
				FirstName: StringPtr("Janet"),
				LastName:  StringPtr("Doensky"),
			},
			err:           &customErrors.InvalidEmail,
			expectedError: &customErrors.InvalidEmail,
		},
		"invalid first name length": {
			user: domain.User{
				Model: gorm.Model{
					ID: 1,
				},
				Username:  StringPtr("janedoe"),
				Password:  StringPtr("P4ssw0rd//"),
				Email:     StringPtr("janedoe2@gmail.com"),
				FirstName: StringPtr("J"),
				LastName:  StringPtr("Doensky"),
			},
			err:           &customErrors.InvalidName,
			expectedError: &customErrors.InvalidName,
		},
		"invalid first name character": {
			user: domain.User{
				Model: gorm.Model{
					ID: 1,
				},
				Username:  StringPtr("janedoe"),
				Password:  StringPtr("P4ssw0rd//"),
				Email:     StringPtr("janedoe2@gmail.com"),
				FirstName: StringPtr("Janette/1"),
				LastName:  StringPtr("Doensky"),
			},
			err:           &customErrors.InvalidName,
			expectedError: &customErrors.InvalidName,
		},
		"invalid last name length": {
			user: domain.User{
				Model: gorm.Model{
					ID: 1,
				},
				Username:  StringPtr("janedoe"),
				Password:  StringPtr("P4ssw0rd//"),
				Email:     StringPtr("janedoe2@gmail.com"),
				FirstName: StringPtr("Janet"),
				LastName:  StringPtr("D"),
			},
			err:           &customErrors.InvalidName,
			expectedError: &customErrors.InvalidName,
		},
		"invalid last name character": {
			user: domain.User{
				Model: gorm.Model{
					ID: 1,
				},
				Username:  StringPtr("janedoe"),
				Password:  StringPtr("P4ssw0rd//"),
				Email:     StringPtr("janedoe2@gmail.com"),
				FirstName: StringPtr("Janet"),
				LastName:  StringPtr("Doesovich@"),
			},
			err:           &customErrors.InvalidName,
			expectedError: &customErrors.InvalidName,
		},

		"existing mail": {
			user: domain.User{
				Model: gorm.Model{
					ID: 1,
				},
				Username:  StringPtr("janedoe"),
				Password:  StringPtr("P4ssw0rd//"),
				Email:     StringPtr("janedoe2@gmail.com"),
				FirstName: StringPtr("Janet"),
				LastName:  StringPtr("Doensky"),
			},
			err:           &customErrors.ExistingEntry,
			expectedError: &customErrors.ExistingEntry,
		},
		"server error": {
			user: domain.User{
				Model: gorm.Model{
					ID: 1,
				},
				Username:  StringPtr("janedoe"),
				Password:  StringPtr("P4ssw0rd//"),
				Email:     StringPtr("janedoe2@gmail.com"),
				FirstName: StringPtr("Janet"),
				LastName:  StringPtr("Doensky"),
			},
			err:           &customErrors.ServerError,
			expectedError: &customErrors.ServerError,
		},
	}

	for testname, subtest := range tests {
		t.Run(testname, func(t *testing.T) {
			mockUserRepository := new(mocks.MockUserRepository)
			userService := user.Service{Repository: mockUserRepository}
			mockUserRepository.On("Update").Return(subtest.err)
			err := userService.Update(subtest.user)

			assert.Equal(t, subtest.expectedError, err)
		})
	}

}
