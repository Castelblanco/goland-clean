package wrappers

import (
	"github.com/Castelblanco/goland-clean/src/modules/tasks/domain/entities"
	"github.com/Castelblanco/goland-clean/src/modules/tasks/infrastructure/storages/sql/models"
)

type TaskWrappers struct {
}

func (w TaskWrappers) DomToDal(task entities.TaskDOM) models.TaskDAL {
	return models.TaskDAL{
		ID:          task.Id,
		Title:       task.Title,
		Description: task.Description,
		Done:        task.Done,
		UserID:      task.User.Id,
	}
}

func (w TaskWrappers) DalToDom(task models.TaskDAL) entities.TaskDOM {
	return entities.TaskDOM{
		Id:          task.ID,
		Title:       task.Title,
		Description: task.Description,
		Done:        task.Done,
	}
}
