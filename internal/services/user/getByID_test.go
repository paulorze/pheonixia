package user_test

import (
	customErrors "phoenixia/errors"
	"phoenixia/internal/domain"
	"phoenixia/internal/repositories/postgreSQL/user/mocks"
	"phoenixia/internal/services/user"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestService_GetByID(t *testing.T) {

	tests := map[string]struct {
		id            string
		user          domain.User
		err           error
		expectedError error
	}{
		"success": {
			id:            "1",
			err:           nil,
			expectedError: nil,
		},
		"invalid id": {
			id:            "",
			err:           &customErrors.InvalidId,
			expectedError: &customErrors.InvalidId,
		},
		"user not found": {
			id:            "1",
			err:           customErrors.UserNotFound,
			expectedError: customErrors.UserNotFound,
		},
		"server error": {
			id:            "1",
			err:           &customErrors.ServerError,
			expectedError: &customErrors.ServerError,
		},
	}

	for testname, subtest := range tests {
		t.Run(testname, func(t *testing.T) {
			mockUserRepository := new(mocks.MockUserRepository)
			userService := user.Service{Repository: mockUserRepository}
			mockUserRepository.On("GetByID").Return(subtest.user, subtest.err)
			user, err := userService.GetByID(subtest.id)

			assert.Equal(t, subtest.user, user)
			assert.Equal(t, subtest.expectedError, err)
		})
	}

}
