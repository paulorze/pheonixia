package user_test

import (
	customErrors "phoenixia/errors"
	"phoenixia/internal/domain"
	"phoenixia/internal/repositories/postgreSQL/user/mocks"
	"phoenixia/internal/services/user"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestService_GetAll(t *testing.T) {

	tests := map[string]struct {
		usersList     []domain.User
		err           error
		expectedError error
	}{
		"success": {
			err:           nil,
			expectedError: nil,
		},
		"server error": {
			err:           &customErrors.ServerError,
			expectedError: &customErrors.ServerError,
		},
	}

	for testname, subtest := range tests {
		t.Run(testname, func(t *testing.T) {
			mockUserRepository := new(mocks.MockUserRepository)
			userService := user.Service{Repository: mockUserRepository}
			mockUserRepository.On("GetAll").Return(subtest.usersList, subtest.err)
			usersList, err := userService.GetAll()

			assert.Equal(t, subtest.usersList, usersList)
			assert.Equal(t, subtest.expectedError, err)
			mockUserRepository.AssertExpectations(t)
		})
	}

}
