package user_test

import (
	"encoding/json"
	"errors"
	"phoenixia/cmd/api/handlers/user"
	customErrors "phoenixia/errors"

	"phoenixia/internal/domain"
	"phoenixia/internal/services/user/mocks"

	"net/http"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

type responseGetAll struct {
	Data  []domain.User `json:"data"`
	Error string        `json:"error"`
}

func TestHandler_GetAll(t *testing.T) {

	tests := map[string]struct {
		mockUsers            []domain.User
		err                  error
		expectedStatus       int
		expectedErrorMessage string
	}{
		"success": {
			mockUsers:            []domain.User{},
			err:                  nil,
			expectedStatus:       fiber.StatusOK,
			expectedErrorMessage: "",
		},
		"server error": {
			mockUsers:            nil,
			err:                  &customErrors.ServerError,
			expectedStatus:       fiber.StatusInternalServerError,
			expectedErrorMessage: customErrors.ServerError.Message,
		},
		"unexpected error": {
			mockUsers:            nil,
			err:                  &customErrors.ExistingEmail,
			expectedStatus:       fiber.StatusBadRequest,
			expectedErrorMessage: "something unexpected happened",
		},
		"internal server error": {
			mockUsers:            nil,
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
			app.Get("/users", handler.GetAll)
			mockUserService.On("GetAll").Return(subtest.mockUsers, subtest.err)

			req, _ := http.NewRequest("GET", "/users", nil)
			resp, respErr := app.Test(req)
			assert.NoError(t, respErr)
			var resBody responseGetAll
			respErr = json.NewDecoder(resp.Body).Decode(&resBody)
			assert.NoError(t, respErr)
			assert.Equal(t, subtest.expectedStatus, resp.StatusCode)
			assert.Equal(t, subtest.mockUsers, resBody.Data)
			assert.Equal(t, subtest.expectedErrorMessage, resBody.Error)
			mockUserService.AssertExpectations(t)
		})
	}
}
