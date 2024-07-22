package services

import "github.com/Castelblanco/goland-clean/src/modules/users/domain/entities"

func (s UsersServices) GetById(id string) (entities.UserDOM, interface{}) {
	return s.Dependencies.Repository.GetById(id)
}
