package services

func (s TaskServices) DeleteOne(id string) interface{} {
	return s.Repository.DeleteOne(id)
}
