package mocks

import (
	"phoenixia/internal/domain"

	"github.com/stretchr/testify/mock"
)

type MockUserRepository struct {
	mock.Mock
}

func (m *MockUserRepository) Create(user domain.User) error {
	args := m.Called()
	return args.Error(0)
}

func (m *MockUserRepository) GetAll() ([]domain.User, error) {
	args := m.Called()
	return args.Get(0).([]domain.User), args.Error(1)
}

func (m *MockUserRepository) GetByID(id string) (domain.User, error) {
	args := m.Called()
	return args.Get(0).(domain.User), args.Error(1)
}

func (m *MockUserRepository) Update(user domain.User) error {
	args := m.Called()
	return args.Error(0)
}

func (m *MockUserRepository) Delete(id string) error {
	args := m.Called()
	return args.Error(0)
}
