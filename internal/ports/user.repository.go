package ports

import "phoenixia/internal/domain"

type UserRepository interface {
	Create(user domain.User) (err error)
	GetAll() (users []domain.User, err error)
	GetByID(id string) (user domain.User, err error)
	Update(updatedUser domain.User) (err error)
	Delete(id string) (err error)
}
