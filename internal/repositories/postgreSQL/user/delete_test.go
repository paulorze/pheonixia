package user_test

import (
	customErrors "phoenixia/errors"
	"phoenixia/internal/domain"
	"phoenixia/internal/repositories/postgreSQL/user"
	"phoenixia/storage"
	"phoenixia/storage/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRepository_Delete(t *testing.T) {

	db, err := storage.NewConnection(mocks.MockPostgresConnect())
	assert.NoError(t, err)
	err = domain.MigrateUser(db)
	assert.NoError(t, err)
	userRepository := user.Repository{
		DB: db,
	}
	firstUser := domain.User{
		Username:  StringPtr("janedoe"),
		Password:  StringPtr("P4ssw0rd//"),
		Email:     StringPtr("janedoe@gmail.com"),
		FirstName: StringPtr("Jane"),
		LastName:  StringPtr("Doe"),
	}
	userRepository.Create(firstUser)
	secondUser := domain.User{
		Username:  StringPtr("johndoe"),
		Password:  StringPtr("P4ssw0rd//"),
		Email:     StringPtr("johndoe@gmail.com"),
		FirstName: StringPtr("John"),
		LastName:  StringPtr("Doe"),
	}
	userRepository.Create(secondUser)

	tests := map[string]struct {
		id            string
		err           error
		expectedError error
	}{
		"success": {
			id:            "2",
			err:           nil,
			expectedError: nil,
		},
		"user not found": {
			id:            "4",
			err:           customErrors.UserNotFound,
			expectedError: customErrors.UserNotFound,
		},
	}

	for testname, subtest := range tests {
		t.Run(testname, func(t *testing.T) {
			err = userRepository.Delete(subtest.id)
			assert.Equal(t, subtest.expectedError, err)
		})
	}

	db.Exec("DROP TABLE users")
}
