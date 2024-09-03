package user

import (
	"errors"
	customErrors "phoenixia/errors"
	"phoenixia/internal/domain"

	"github.com/jackc/pgx/v5/pgconn"
)

func (repository *Repository) Update(user domain.User) (err error) {
	err = repository.DB.Save(&user).Error
	if err != nil {
		var perr *pgconn.PgError
		errors.As(err, &perr)
		if perr.Code == "23505" {
			err = &customErrors.ExistingEntry
			return
		}
		err = &customErrors.ServerError
		return
	}
	return
}
