package services

import "github.com/Castelblanco/goland-clean/src/modules/users/domain/entities"

func (s UsersServices) UpdateOne(user entities.UserDOM) (entities.UserDOM, interface{}) {
	return s.Dependencies.Repository.UpdateOne(user)
}
