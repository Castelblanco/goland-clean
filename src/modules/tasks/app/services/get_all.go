package services

import "github.com/Castelblanco/goland-clean/src/modules/tasks/domain/entities"

func (s TaskServices) GetAll() ([]entities.TaskDOM, interface{}) {
	return s.Repository.GetAll()
}
