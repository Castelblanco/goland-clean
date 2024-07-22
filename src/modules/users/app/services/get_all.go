package services

import (
	"github.com/Castelblanco/goland-clean/src/modules/users/domain/entities"
)

func (s UsersServices) GetAll() ([]entities.UserDOM, interface{}) {
	return s.Dependencies.Repository.GetAll()
}
