package user_test

import (
	"errors"
	"phoenixia/cmd/api/handlers/user"
	customErrors "phoenixia/errors"

	"encoding/json"
	"phoenixia/internal/domain"
	"phoenixia/internal/services/user/mocks"

	"net/http"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

type responseGetByID struct {
	Data  domain.User `json:"data"`
	Error string      `json:"error"`
}

func TestHandler_GetByID(t *testing.T) {

	tests := map[string]struct {
		idParam              string
		mockUser             domain.User
		err                  error
		expectedStatus       int
		expectedErrorMessage string
	}{
		"success": {
			idParam:              "1",
			err:                  nil,
			expectedStatus:       fiber.StatusOK,
			expectedErrorMessage: "",
		},
		"user not found": {
			idParam:              "1",
			err:                  customErrors.UserNotFound,
			expectedStatus:       fiber.StatusNotFound,
			expectedErrorMessage: customErrors.UserNotFound.Message,
		},
		"server error": {
			idParam:              "1",
			err:                  &customErrors.ServerError,
			expectedStatus:       fiber.StatusInternalServerError,
			expectedErrorMessage: customErrors.ServerError.Message,
		},
		"unexpected error": {
			idParam:              "1",
			err:                  &customErrors.ExistingEmail,
			expectedStatus:       fiber.StatusBadRequest,
			expectedErrorMessage: "something unexpected happened",
		},
		"internal server error": {
			idParam:              "1",
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
			app.Get("/users/:id", handler.GetByID)
			mockUserService.On("GetByID").Return(subtest.mockUser, subtest.err)

			endpoint := "/users/" + subtest.idParam
			req, _ := http.NewRequest("GET", endpoint, nil)
			resp, respErr := app.Test(req)
			assert.NoError(t, respErr)
			var resBody responseGetByID
			respErr = json.NewDecoder(resp.Body).Decode(&resBody)
			assert.NoError(t, respErr)
			assert.Equal(t, subtest.expectedStatus, resp.StatusCode)
			assert.Equal(t, subtest.mockUser, resBody.Data)
			assert.Equal(t, subtest.expectedErrorMessage, resBody.Error)
			mockUserService.AssertExpectations(t)
		})
	}

}
