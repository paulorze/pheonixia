package mocks

import (
	"phoenixia/internal/domain"

	"github.com/stretchr/testify/mock"
)

type MockUserService struct {
	mock.Mock
}

func (m *MockUserService) Create(user domain.User) error {
	args := m.Called()
	return args.Error(0)
}

func (m *MockUserService) GetAll() ([]domain.User, error) {
	args := m.Called()
	return args.Get(0).([]domain.User), args.Error(1)
}

func (m *MockUserService) GetByID(id string) (domain.User, error) {
	args := m.Called()
	return args.Get(0).(domain.User), args.Error(1)
}

func (m *MockUserService) Update(user domain.User) error {
	args := m.Called()
	return args.Error(0)
}

func (m *MockUserService) Delete(id string) error {
	args := m.Called()
	return args.Error(0)
}
