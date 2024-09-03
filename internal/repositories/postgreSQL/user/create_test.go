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

func StringPtr(str string) *string {
	return &str
}

func TestRepository_Create(t *testing.T) {

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

	tests := map[string]struct {
		user          domain.User
		err           error
		expectedError error
	}{
		"success": {
			user: domain.User{
				Username:  StringPtr("johndoe"),
				Password:  StringPtr("P4ssw0rd//"),
				Email:     StringPtr("johndoe@gmail.com"),
				FirstName: StringPtr("John"),
				LastName:  StringPtr("Doe"),
			},
			err:           nil,
			expectedError: nil,
		},
		"existing username": {
			user: domain.User{
				Username:  StringPtr("janedoe"),
				Password:  StringPtr("P4ssw0rd//"),
				Email:     StringPtr("janedo3@gmail.com"),
				FirstName: StringPtr("Jane"),
				LastName:  StringPtr("Doe"),
			},
			err:           &customErrors.ExistingEntry,
			expectedError: &customErrors.ExistingEntry,
		},
		"existing email": {
			user: domain.User{
				Username:  StringPtr("janedo3"),
				Password:  StringPtr("P4ssw0rd//"),
				Email:     StringPtr("janedoe@gmail.com"),
				FirstName: StringPtr("Jane"),
				LastName:  StringPtr("Doe"),
			},
			err:           &customErrors.ExistingEntry,
			expectedError: &customErrors.ExistingEntry,
		},
	}

	for testname, subtest := range tests {
		t.Run(testname, func(t *testing.T) {
			err = userRepository.Create(subtest.user)
			assert.Equal(t, subtest.expectedError, err)
		})
	}

	db.Exec("DROP TABLE users")
}
