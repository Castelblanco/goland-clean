package repository

import (
	"github.com/Castelblanco/goland-clean/src/modules/users/domain/entities"
)

type UserRepository interface {
	GetAll() ([]entities.UserDOM, interface{})
	GetById(id string) (entities.UserDOM, interface{})
	CreateOne(user entities.UserDOM) (entities.UserDOM, interface{})
	UpdateOne(user entities.UserDOM) (entities.UserDOM, interface{})
	DeleteOne(id string) interface{}
}
