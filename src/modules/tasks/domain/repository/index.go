package repository

import "github.com/Castelblanco/goland-clean/src/modules/tasks/domain/entities"

type TaskRepository interface {
	GetAll() ([]entities.TaskDOM, interface{})
	GetById(id string) (entities.TaskDOM, interface{})
	CreateOne(task entities.TaskDOM) (entities.TaskDOM, interface{})
	UpdateOne(task entities.TaskDOM) (entities.TaskDOM, interface{})
	DeleteOne(id string) interface{}
}
