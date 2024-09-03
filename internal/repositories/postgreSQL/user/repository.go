package user

import (
	"phoenixia/internal/ports"

	"gorm.io/gorm"
)

var _ ports.UserRepository = &Repository{}

type Repository struct {
	DB *gorm.DB
}
