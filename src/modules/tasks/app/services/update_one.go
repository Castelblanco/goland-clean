package services

import "github.com/Castelblanco/goland-clean/src/modules/tasks/domain/entities"

func (s TaskServices) UpdateOne(task entities.TaskDOM) (entities.TaskDOM, interface{}) {
	return s.Repository.UpdateOne(task)
}
