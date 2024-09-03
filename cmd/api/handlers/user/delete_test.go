package user_test

import (
	"encoding/json"
	"errors"
	"net/http"
	"phoenixia/cmd/api/handlers/user"
	customErrors "phoenixia/errors"
	"phoenixia/internal/services/user/mocks"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

type responseWithMessage struct {
	Message string `json:"message"`
	Error   string `json:"error"`
}

func TestHandler_Delete(t *testing.T) {

	tests := map[string]struct {
		idParam              string
		err                  error
		expectedStatus       int
		expectedMessage      string
		expectedErrorMessage string
	}{
		"success": {
			idParam:              "1",
			err:                  nil,
			expectedStatus:       fiber.StatusOK,
			expectedMessage:      "user deleted successfully",
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
			app.Delete("/users/:id", handler.Delete)
			mockUserService.On("Delete").Return(subtest.err)

			endpoint := "/users/" + subtest.idParam
			req, _ := http.NewRequest("DELETE", endpoint, nil)
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
