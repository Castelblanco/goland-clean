package services

import "github.com/Castelblanco/goland-clean/src/modules/tasks/domain/repository"

type Dependencies struct {
	Repository repository.TaskRepository
	CreatedId  func() string
}

type TaskServices struct {
	Dependencies
}
