package user_test

import (
	"phoenixia/internal/domain"
	"phoenixia/internal/repositories/postgreSQL/user"
	"phoenixia/storage"
	"phoenixia/storage/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRepository_GetAll(t *testing.T) {

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
		userList      []domain.User
		err           error
		expectedError error
	}{
		"success": {
			userList:      []domain.User{firstUser, secondUser},
			err:           nil,
			expectedError: nil,
		},
	}

	for testname, subtest := range tests {
		t.Run(testname, func(t *testing.T) {
			userList, err := userRepository.GetAll()
			assert.Equal(t, subtest.userList[0].Username, userList[0].Username)
			assert.Equal(t, subtest.expectedError, err)
		})
	}

	db.Exec("DROP TABLE users")
}
