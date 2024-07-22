package success

import "reflect"

type ListResponse struct {
	Items  []interface{} `json:"items"`
	Count  int           `json:"count"`
	Status int           `json:"status"`
}

func NewListResponse(items interface{}, count int, status int) ListResponse {
	return ListResponse{
		Items:  makeItems(items),
		Count:  count,
		Status: status,
	}
}

func makeItems(items interface{}) []interface{} {
	v := reflect.ValueOf(items)
	if v.Kind() != reflect.Slice {
		return nil
	}

	list := make([]interface{}, v.Len())
	for i := 0; i < v.Len(); i++ {
		list[i] = v.Index(i).Interface()
	}
	return list
}
