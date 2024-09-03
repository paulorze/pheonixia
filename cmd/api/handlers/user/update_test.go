package user_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"phoenixia/cmd/api/handlers/user"
	customErrors "phoenixia/errors"
	"phoenixia/internal/domain"
	"phoenixia/internal/services/user/mocks"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func StringPtr(str string) *string {
	return &str
}

func TestHandler_Update(t *testing.T) {

	userUpdateTemplate := domain.User{
		Email:     StringPtr("carlosmerca@prueba.com"),
		FirstName: StringPtr("Carlos"),
		LastName:  StringPtr("Merca"),
	}

	tests := map[string]struct {
		idParam              string
		user                 domain.User
		err                  error
		expectedStatus       int
		expectedMessage      string
		expectedErrorMessage string
	}{
		"success": {
			idParam:              "1",
			user:                 userUpdateTemplate,
			err:                  nil,
			expectedStatus:       fiber.StatusOK,
			expectedMessage:      "user updated successfully",
			expectedErrorMessage: "",
		},
		"user not found": {
			idParam:              "1",
			user:                 userUpdateTemplate,
			err:                  customErrors.UserNotFound,
			expectedStatus:       fiber.StatusNotFound,
			expectedErrorMessage: customErrors.UserNotFound.Message,
		},
		"existing mail": {
			idParam:              "1",
			user:                 userUpdateTemplate,
			err:                  &customErrors.ExistingEntry,
			expectedStatus:       fiber.StatusConflict,
			expectedErrorMessage: customErrors.ExistingEntry.Message,
		},
		"invalid data": {
			idParam:              "1",
			user:                 userUpdateTemplate,
			err:                  &customErrors.InvalidEmail,
			expectedStatus:       fiber.StatusUnprocessableEntity,
			expectedErrorMessage: customErrors.InvalidEmail.Message,
		},
		"server error": {
			idParam:              "1",
			user:                 userUpdateTemplate,
			err:                  &customErrors.ServerError,
			expectedStatus:       fiber.StatusInternalServerError,
			expectedErrorMessage: customErrors.ServerError.Message,
		},
		"unexpected error": {
			idParam:              "1",
			user:                 userUpdateTemplate,
			err:                  &customErrors.MockError,
			expectedStatus:       fiber.StatusBadRequest,
			expectedErrorMessage: "something unexpected happened",
		},
		"internal server error": {
			idParam:              "1",
			user:                 userUpdateTemplate,
			err:                  errors.New("RANDOM ERROR"),
			expectedStatus:       fiber.StatusInternalServerError,
			expectedErrorMessage: "internal server error",
		},
	}

	for testname, subtest := range tests {
		t.Run(testname, func(t *testing.T) {
			app := fiber.New()
			mockUserService := new(mocks.MockUserService)
			handler := user.Handler{UserService: mockUserService}
			app.Put("/users", handler.Update)
			mockUserService.On("Update").Return(subtest.err)

			body, err := json.Marshal(subtest.user)
			assert.NoError(t, err)
			req, err := http.NewRequest("PUT", "/users", bytes.NewBuffer(body))
			req.Header.Set("Content-Type", "application/json")
			assert.NoError(t, err)
			resp, respErr := app.Test(req)
			assert.NoError(t, respErr)
			var resBody responseWithMessage
			respErr = json.NewDecoder(resp.Body).Decode(&resBody)
			assert.NoError(t, respErr)
			assert.Equal(t, subtest.expectedStatus, resp.StatusCode)
			assert.Equal(t, subtest.expectedMessage, resBody.Message)
			assert.Equal(t, subtest.expectedErrorMessage, resBody.Error)
			mockUserService.AssertExpectations(t)
		})
	}

}
