package implementation

import (
	"fmt"

	"github.com/Castelblanco/goland-clean/src/app/db"
	"github.com/Castelblanco/goland-clean/src/modules/common/responses/errors"
	"github.com/Castelblanco/goland-clean/src/modules/tasks/domain/entities"
	"github.com/Castelblanco/goland-clean/src/modules/tasks/infrastructure/storages/sql/models"
	"github.com/Castelblanco/goland-clean/src/modules/tasks/infrastructure/storages/sql/wrappers"
)

var DEFAULT_METADATA = map[string]interface{}{}

type TaskSqlRepository struct {
	Wrappers wrappers.TaskWrappers
}

func (r TaskSqlRepository) GetAll() ([]entities.TaskDOM, interface{}) {
	dal := []models.TaskDAL{}
	result := db.DB.Find(&dal)

	if result.Error != nil {
		return []entities.TaskDOM{}, errors.NewStorageError(result.Error.Error(), DEFAULT_METADATA)
	}

	tasks := []entities.TaskDOM{}
	for _, task := range dal {
		tasks = append(tasks, r.Wrappers.DalToDom(task))
	}
	return tasks, nil
}

func (r TaskSqlRepository) GetById(id string) (entities.TaskDOM, interface{}) {
	dal := models.TaskDAL{}
	result := db.DB.First(&dal, models.TaskDAL{
		ID: id,
	})

	if result.Error != nil {
		return entities.TaskDOM{}, errors.NewStorageError(result.Error.Error(), DEFAULT_METADATA)
	}
	return r.Wrappers.DalToDom(dal), nil
}

func (r TaskSqlRepository) CreateOne(task entities.TaskDOM) (entities.TaskDOM, interface{}) {
	dal := r.Wrappers.DomToDal(task)
	fmt.Println(dal)
	fmt.Println(dal.UserID)
	result := db.DB.Create(&dal)

	if result.Error != nil {
		return entities.TaskDOM{}, errors.NewStorageError(result.Error.Error(), DEFAULT_METADATA)
	}
	return r.Wrappers.DalToDom(dal), nil
}

func (r TaskSqlRepository) UpdateOne(task entities.TaskDOM) (entities.TaskDOM, interface{}) {
	dal := r.Wrappers.DomToDal(task)
	result := db.DB.Save(&dal)

	if result.Error != nil {
		return entities.TaskDOM{}, errors.NewStorageError(result.Error.Error(), DEFAULT_METADATA)
	}
	return r.Wrappers.DalToDom(dal), nil
}

func (r TaskSqlRepository) DeleteOne(id string) interface{} {
	result := db.DB.Unscoped().Delete(&models.TaskDAL{
		ID: id,
	})

	if result.Error != nil {
		return errors.NewStorageError(result.Error.Error(), DEFAULT_METADATA)
	}

	if result.RowsAffected == 0 {
		return errors.NewStorageError("record not found", DEFAULT_METADATA)
	}
	return nil
}
