package implementation

import (
	"fmt"

	"github.com/Castelblanco/goland-clean/src/app/db"
	"github.com/Castelblanco/goland-clean/src/modules/common/responses/errors"
	"github.com/Castelblanco/goland-clean/src/modules/users/domain/entities"
	"github.com/Castelblanco/goland-clean/src/modules/users/infrastructure/storages/sql/models"
	"github.com/Castelblanco/goland-clean/src/modules/users/infrastructure/storages/sql/wrapper"
)

var DEFAULT_METADATA = map[string]interface{}{}

type UsersSqlRepository struct {
	Wrapper wrapper.UsersWrapper
}

func (r UsersSqlRepository) GetAll() ([]entities.UserDOM, interface{}) {
	var dal []models.UserDAL
	result := db.DB.Find(&dal)

	if result.Error != nil {
		return []entities.UserDOM{}, errors.NewErrorBadRequest(result.Error.Error(), DEFAULT_METADATA)
	}

	var users []entities.UserDOM
	for _, user := range dal {
		users = append(users, r.Wrapper.DalToDom(user))
	}
	return users, nil
}

func (r UsersSqlRepository) GetById(id string) (entities.UserDOM, interface{}) {
	var dal models.UserDAL
	result := db.DB.First(&dal, models.UserDAL{
		ID: id,
	})

	if result.Error != nil {
		return entities.UserDOM{}, errors.NewStorageError(result.Error.Error(), DEFAULT_METADATA)
	}

	err := db.DB.Model(&dal).Association("TASKS").Find(&dal.Tasks)

	fmt.Println(err)
	return r.Wrapper.DalToDom(dal), nil
}

func (r UsersSqlRepository) CreateOne(user entities.UserDOM) (entities.UserDOM, interface{}) {
	dal := r.Wrapper.DomToDal(user)
	result := db.DB.Create(&dal)

	if result.Error != nil {
		return entities.UserDOM{}, errors.NewStorageError(result.Error.Error(), DEFAULT_METADATA)
	}
	return r.Wrapper.DalToDom(dal), nil
}

func (r UsersSqlRepository) UpdateOne(user entities.UserDOM) (entities.UserDOM, interface{}) {
	dal := r.Wrapper.DomToDal(user)
	result := db.DB.Save(&dal)

	if result.Error != nil {
		return entities.UserDOM{}, errors.NewStorageError(result.Error.Error(), DEFAULT_METADATA)
	}
	return r.Wrapper.DalToDom(dal), nil
}

func (r UsersSqlRepository) DeleteOne(id string) interface{} {
	result := db.DB.Unscoped().Delete(&models.UserDAL{
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
