package services

import "github.com/Castelblanco/goland-clean/src/modules/users/domain/entities"

func (s UsersServices) CreateOne(user entities.UserDOM) (entities.UserDOM, interface{}) {
	user.Id = s.Dependencies.CreatedId()
	return s.Dependencies.Repository.CreateOne(user)
}
