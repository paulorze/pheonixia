package user

import (
	"errors"
	"fmt"
	customErrors "phoenixia/errors"
	"phoenixia/internal/domain"

	"github.com/jackc/pgx/v5/pgconn"
)

func (repository *Repository) Create(user domain.User) (err error) {
	err = repository.DB.Create(&user).Error
	if err != nil {
		var perr *pgconn.PgError
		errors.As(err, &perr)
		fmt.Println(perr.Code)
		if perr.Code == "23505" {
			err = &customErrors.ExistingEntry
			return
		}
		err = &customErrors.ServerError
		return
	}
	return
}
