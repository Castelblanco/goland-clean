package entities

type UserDOM struct {
	Id        string        `json:"_id"`
	FirstName string        `json:"first_name"`
	LastName  string        `json:"last_name"`
	Email     string        `json:"email"`
	Tasks     []UserTaskDOM `json:"tasks"`
}

type UserTaskDOM struct {
	Id          string `json:"_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Done        bool   `json:"done"`
}
