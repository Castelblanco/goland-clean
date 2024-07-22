package services

func (s UsersServices) DeleteOne(id string) interface{} {
	return s.Dependencies.Repository.DeleteOne(id)
}
