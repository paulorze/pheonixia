package user

import "phoenixia/internal/ports"

type Handler struct {
	UserService ports.UserService
}
