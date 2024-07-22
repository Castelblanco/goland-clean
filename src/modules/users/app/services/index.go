package services

import "github.com/Castelblanco/goland-clean/src/modules/users/domain/repository"

type Dependencies struct {
	Repository repository.UserRepository
	CreatedId  func() string
}

type UsersServices struct {
	Dependencies
}
