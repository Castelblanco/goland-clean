package services

import "github.com/Castelblanco/goland-clean/src/modules/tasks/domain/entities"

func (s TaskServices) GetById(id string) (entities.TaskDOM, interface{}) {
	return s.Repository.GetById(id)
}
