package wrapper

import (
	"github.com/Castelblanco/goland-clean/src/modules/users/domain/entities"
	"github.com/Castelblanco/goland-clean/src/modules/users/infrastructure/storages/sql/models"
)

type UsersWrapper struct {
}

func (w UsersWrapper) DomToDal(user entities.UserDOM) models.UserDAL {
	tasks := []models.UserTaskDAL{}
	for _, task := range user.Tasks {
		tasks = append(tasks, models.UserTaskDAL{
			ID:          task.Id,
			Title:       task.Title,
			Description: task.Description,
			Done:        task.Done,
			UserID:      user.Id,
		})
	}
	return models.UserDAL{
		ID:        user.Id,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		Tasks:     tasks,
	}
}

func (w UsersWrapper) DalToDom(user models.UserDAL) entities.UserDOM {
	tasks := []entities.UserTaskDOM{}

	for _, task := range user.Tasks {
		tasks = append(tasks, entities.UserTaskDOM{
			Id:          task.ID,
			Title:       task.Title,
			Description: task.Description,
			Done:        task.Done,
		})
	}
	return entities.UserDOM{
		Id:        user.ID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		Tasks:     tasks,
	}
}
