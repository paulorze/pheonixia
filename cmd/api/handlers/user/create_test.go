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

func TestHandler_Create(t *testing.T) {

	userCreateTemplate := domain.User{
		Username:  StringPtr("carlosmerca"),
		Password:  StringPtr("Probando123//"),
		Email:     StringPtr("carlosmerca@prueba.com"),
		FirstName: StringPtr("Carlos"),
		LastName:  StringPtr("Merca"),
	}

	tests := map[string]struct {
		user                 domain.User
		err                  error
		expectedStatus       int
		expectedMessage      string
		expectedErrorMessage string
	}{
		"success": {
			user:                 userCreateTemplate,
			err:                  nil,
			expectedStatus:       fiber.StatusCreated,
			expectedMessage:      "user created successfully",
			expectedErrorMessage: "",
		},
		"incomplete register fields": {
			user:                 userCreateTemplate,
			err:                  &customErrors.IncompleteRegisterFields,
			expectedStatus:       fiber.StatusBadRequest,
			expectedErrorMessage: customErrors.IncompleteRegisterFields.Message,
		},
		"existing mail": {
			user:                 userCreateTemplate,
			err:                  &customErrors.ExistingEntry,
			expectedStatus:       fiber.StatusConflict,
			expectedErrorMessage: customErrors.ExistingEntry.Message,
		},
		"invalid data": {
			user:                 userCreateTemplate,
			err:                  &customErrors.InvalidEmail,
			expectedStatus:       fiber.StatusUnprocessableEntity,
			expectedErrorMessage: customErrors.InvalidEmail.Message,
		},
		"server error": {
			user:                 userCreateTemplate,
			err:                  &customErrors.ServerError,
			expectedStatus:       fiber.StatusInternalServerError,
			expectedErrorMessage: customErrors.ServerError.Message,
		},
		"unexpected error": {
			user:                 userCreateTemplate,
			err:                  &customErrors.MockError,
			expectedStatus:       fiber.StatusBadRequest,
			expectedErrorMessage: "something unexpected happened",
		},
		"internal server error": {
			user:                 userCreateTemplate,
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
			app.Post("/users", handler.Create)
			mockUserService.On("Create").Return(subtest.err)

			body, err := json.Marshal(subtest.user)
			assert.NoError(t, err)
			req, err := http.NewRequest("POST", "/users", bytes.NewBuffer(body))
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
