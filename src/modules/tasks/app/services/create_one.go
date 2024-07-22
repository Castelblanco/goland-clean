package services

import "github.com/Castelblanco/goland-clean/src/modules/tasks/domain/entities"

func (s TaskServices) CreateOne(task entities.TaskDOM) (entities.TaskDOM, interface{}) {
	task.Id = s.CreatedId()
	return s.Repository.CreateOne(task)
}
