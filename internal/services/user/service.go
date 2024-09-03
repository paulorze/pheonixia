package user

import "phoenixia/internal/ports"

var _ ports.UserService = &Service{}

type Service struct {
	Repository ports.UserRepository
}
