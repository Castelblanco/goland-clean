package entities

type TaskDOM struct {
	Id          string      `json:"_id"`
	Title       string      `json:"title"`
	Description string      `json:"description"`
	Done        bool        `json:"done"`
	User        TaskUserDOM `json:"user"`
}

type TaskUserDOM struct {
	Id        string `json:"_id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
}
