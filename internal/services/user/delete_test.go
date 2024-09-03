package user_test

import (
	customErrors "phoenixia/errors"
	"phoenixia/internal/repositories/postgreSQL/user/mocks"
	"phoenixia/internal/services/user"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestService_Delete(t *testing.T) {

	tests := map[string]struct {
		id            string
		err           error
		expectedError error
	}{
		"success": {
			id:            "1",
			err:           nil,
			expectedError: nil,
		},
		"missing id": {
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
			mockUserRepository.On("Delete").Return(subtest.err)
			err := userService.Delete(subtest.id)

			assert.Equal(t, subtest.expectedError, err)
			// mockUserRepository.AssertExpectations(t)
		})
	}

}
