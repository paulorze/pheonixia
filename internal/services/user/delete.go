package user

import customErrors "phoenixia/errors"

func (service *Service) Delete(id string) (err error) {
	if id == "" {
		err = &customErrors.InvalidId
		return
	}
	err = service.Repository.Delete(id)
	return
}
